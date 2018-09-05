package client

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeLogger struct {
	output []string
}

func (f *fakeLogger) Debug(s string, args ...interface{}) {
	f.output = append(f.output, s)
}

var _ = Describe("clientFilteringLogger", func() {
	It("filters out the Authorization value in the Headers field", func() {
		fakeLog := &fakeLogger{output: []string{}}

		println("hello")
		logger := newFilteringLogger(fakeLog)
		logger("Data1Authorization:[123456]Data2")
		Expect(fakeLog.output).To(Equal([]string{"[machinebox/graphql] Data1Authorization:[ TOKEN HIDDEN ]Data2"}))
	})

})

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}
