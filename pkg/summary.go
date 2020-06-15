package pkg

import (
	"context"
	"sort"
	"strconv"
	"strings"

	"github.com/cybozu-go/topolvm"
	topolvmv1 "github.com/cybozu-go/topolvm/api/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Summary struct {
	Nodes         []*NodeSummary
	DeviceClasses []string
}

type NodeSummary struct {
	Name       string
	Capacities map[string]uint64
	Used       map[string]uint64
}

func Summarize(cli client.Client) (*Summary, error) {
	nodes := corev1.NodeList{}
	err := cli.List(context.Background(), &nodes, &client.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodeSummaries := make(map[string]*NodeSummary)
	deviceClasses := make(map[string]struct{})
	for _, node := range nodes.Items {
		nodeSummary := NodeSummary{
			Name: node.Name,
		}
		nodeSummary.Capacities = make(map[string]uint64)
		nodeSummary.Used = make(map[string]uint64)

		for k, v := range node.Annotations {
			if strings.HasPrefix(k, topolvm.CapacityKeyPrefix) {
				dc := k[len(topolvm.CapacityKeyPrefix):]
				deviceClasses[dc] = struct{}{}
				bytes, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					return nil, err
				}
				nodeSummary.Capacities[dc] = bytes
			}
		}
		nodeSummaries[node.Name] = &nodeSummary
	}

	var lvlist topolvmv1.LogicalVolumeList
	err = cli.List(context.Background(), &lvlist, &client.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, lv := range lvlist.Items {
		dc := lv.Spec.DeviceClass
		if dc == topolvm.DefaultDeviceClassName {
			dc = topolvm.DefaultDeviceClassAnnotationName
		}
		nodeSummaries[lv.Spec.NodeName].Used[dc] += uint64(lv.Status.CurrentSize.Value())
	}

	nodeNames := make([]string, len(nodeSummaries))
	index := 0
	for k := range nodeSummaries {
		nodeNames[index] = k
		index++
	}
	sort.Strings(nodeNames)
	result := Summary{}
	result.Nodes = make([]*NodeSummary, len(nodeNames))
	for i, name := range nodeNames {
		result.Nodes[i] = nodeSummaries[name]
	}
	result.DeviceClasses = make([]string, len(deviceClasses))
	index = 0
	for dc := range deviceClasses {
		result.DeviceClasses[index] = dc
		index++
	}
	sort.Strings(result.DeviceClasses)

	return &result, nil
}
