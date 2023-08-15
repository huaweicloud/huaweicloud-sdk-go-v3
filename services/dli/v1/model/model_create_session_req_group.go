package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSessionReqGroup 创建会话请求参数groups的元素。
type CreateSessionReqGroup struct {

	// 用户组名称。
	Name *string `json:"name,omitempty"`

	// 用户组资源。
	Resources *[]CreateSessionReqResource `json:"resources,omitempty"`
}

func (o CreateSessionReqGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSessionReqGroup struct{}"
	}

	return strings.Join([]string{"CreateSessionReqGroup", string(data)}, " ")
}
