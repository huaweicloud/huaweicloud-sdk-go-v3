package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeNodePoolResponse Response Object
type UpgradeNodePoolResponse struct {

	// Job ID returned after the job is delivered. The job ID can be used to query the job execution status.
	Jobid          *string `json:"jobid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeNodePoolResponse struct{}"
	}

	return strings.Join([]string{"UpgradeNodePoolResponse", string(data)}, " ")
}
