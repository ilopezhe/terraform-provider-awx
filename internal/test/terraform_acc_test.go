// +build !full

package test

import (
	"strconv"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAcc(t *testing.T) {
	terraformOptions := &terraform.Options{
	    TerraformDir: "../../examples/full",
	    // Vars: map[string]interface{}{
	    //     // additional variables
	    // },
	    // EnvVars: map[string]string{
	    //     "TF_LOG":      "DEBUG",
	    //     "TF_LOG_PATH": "../../terraform.log",
	    // },
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "inventory")
	nr, err := strconv.Atoi(output)
	if err != nil {
		t.Logf("Inventory id is not a number")
		t.Fail()
	}
	assert.Greater(t, nr, 0)
}
