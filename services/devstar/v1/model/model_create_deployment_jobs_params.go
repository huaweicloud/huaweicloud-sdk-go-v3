package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDeploymentJobsParams struct {
	Function *FgsDeploymentJobsParam `json:"function,omitempty"`
}

func (o CreateDeploymentJobsParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentJobsParams struct{}"
	}

	return strings.Join([]string{"CreateDeploymentJobsParams", string(data)}, " ")
}
