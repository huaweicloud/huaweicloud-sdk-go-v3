package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceErrorListResponse struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ResourceErrorListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceErrorListResponse struct{}"
	}

	return strings.Join([]string{"ResourceErrorListResponse", string(data)}, " ")
}
