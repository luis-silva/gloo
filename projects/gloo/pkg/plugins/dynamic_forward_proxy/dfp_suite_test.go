<<<<<<<< HEAD:projects/gloo/pkg/plugins/enterprise_warning/enterprise_warning_suite_test.go
package enterprise_warning_test
========
package dynamic_forward_proxy_test
>>>>>>>> master:projects/gloo/pkg/plugins/dynamic_forward_proxy/dfp_suite_test.go

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

<<<<<<<< HEAD:projects/gloo/pkg/plugins/enterprise_warning/enterprise_warning_suite_test.go
func TestEnterpriseWarning(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "EnterpriseWarning Suite", []Reporter{junitReporter})
========
func TestDfp(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "DynamicForwardProxy Suite", []Reporter{junitReporter})
>>>>>>>> master:projects/gloo/pkg/plugins/dynamic_forward_proxy/dfp_suite_test.go
}
