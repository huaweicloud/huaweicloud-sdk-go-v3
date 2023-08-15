package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CciDeploymentJobsParam struct {

	// 软件包或者镜像地址
	Image *string `json:"image,omitempty"`
}

func (o CciDeploymentJobsParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CciDeploymentJobsParam struct{}"
	}

	return strings.Join([]string{"CciDeploymentJobsParam", string(data)}, " ")
}
