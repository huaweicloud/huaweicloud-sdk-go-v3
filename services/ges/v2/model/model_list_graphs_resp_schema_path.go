package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListGraphsRespSchemaPath struct {

	// 导入OBS文件对应的jobId。
	JobId *string `json:"job_id,omitempty"`

	// OBS存储路径，不包含OBS Endpoint。
	Path *string `json:"path,omitempty"`

	// OBS文件导入状态。  - success：完全导入成功 - partiallyFailed：部分失败 - failed：完全导入失败
	Status *string `json:"status,omitempty"`
}

func (o ListGraphsRespSchemaPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphsRespSchemaPath struct{}"
	}

	return strings.Join([]string{"ListGraphsRespSchemaPath", string(data)}, " ")
}
