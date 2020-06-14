package pkg

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	topolvmv1 "github.com/cybozu-go/topolvm/api/v1"
)

var units = []string{
	"Bytes",
	"KiB",
	"MiB",
	"GiB",
	"TiB",
	"PiB",
	"EiB",
}

func FormatBytes(bytes uint64) string {
	count := 0
	num := float64(bytes)
	for ; ; count++ {
		if num < 1024 {
			break
		}
		num /= 1024
	}
	return fmt.Sprintf("%.1f%s", num, units[count])
}

func PrintLVList(lvlist *topolvmv1.LogicalVolumeList) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	columnNames := []string{"NAME", "NODE", "SIZE", "VOLUME_ID"}
	_, err := fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t"))
	if err != nil {
		return err
	}
	for _, lv := range lvlist.Items {
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", lv.Name, lv.Spec.NodeName, lv.Status.CurrentSize, lv.Status.VolumeID)
		if err != nil {
			return err
		}
	}
	return w.Flush()
}
