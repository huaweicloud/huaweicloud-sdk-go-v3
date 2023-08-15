package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResInstanceInfoRequest Request Object
type ShowResInstanceInfoRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 禁用/启用
	Action string `json:"action"`

	Body *ListEnterpriseResourceRequestBody `json:"body,omitempty"`
}

func (o ShowResInstanceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResInstanceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowResInstanceInfoRequest", string(data)}, " ")
}
