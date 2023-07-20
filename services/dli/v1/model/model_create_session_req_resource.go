package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSessionReqResource 创建会话请求参数resources的元素。
type CreateSessionReqResource struct {

	// 资源名称。
	Name *string `json:"name,omitempty"`

	// 资源类型。
	Type *string `json:"type,omitempty"`
}

func (o CreateSessionReqResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSessionReqResource struct{}"
	}

	return strings.Join([]string{"CreateSessionReqResource", string(data)}, " ")
}
