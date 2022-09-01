package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FgsDeploymentJobsParam struct {

	// 版本包id,使用历史版本包部署时需要
	FileId *string `json:"file_id,omitempty" xml:"file_id"`

	// 函数入口
	Handler *string `json:"handler,omitempty" xml:"handler"`
}

func (o FgsDeploymentJobsParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FgsDeploymentJobsParam struct{}"
	}

	return strings.Join([]string{"FgsDeploymentJobsParam", string(data)}, " ")
}
