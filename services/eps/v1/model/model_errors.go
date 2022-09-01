package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 错误列表
type Errors struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// ProjectID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type"`
}

func (o Errors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Errors struct{}"
	}

	return strings.Join([]string{"Errors", string(data)}, " ")
}
