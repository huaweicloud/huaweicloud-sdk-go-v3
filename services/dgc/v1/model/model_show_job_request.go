package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobRequest Request Object
type ShowJobRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 作业名称.
	JobName string `json:"job_name"`

	// 作业版本号，若传入版本号，则查询指定版本号的作业；若不传入，则查询最新的版本作业.
	Version *int32 `json:"version,omitempty"`

	// 返回下游依赖当前作业的作业，只返回第一层。
	Dependencies *bool `json:"dependencies,omitempty"`
}

func (o ShowJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
