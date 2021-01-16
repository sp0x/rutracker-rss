package indexer

import (
	"github.com/onsi/gomega"
	"testing"
)

func Test_invokeFilter_ShouldHandleDateparse(t *testing.T) {
	g := gomega.NewWithT(t)
	format := "200601021504"
	//"2006-01-02T15:04:05Z07:00"
	timeTS := "202007021120"

	result, err := invokeFilter("dateparse", format, timeTS)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result).ToNot(gomega.BeNil())
}
