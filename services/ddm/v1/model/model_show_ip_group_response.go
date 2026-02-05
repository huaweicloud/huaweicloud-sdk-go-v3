package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpGroupResponse Response Object
type ShowIpGroupResponse struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **参数范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  组ID。此参数是参数组的唯一标识。  **参数范围**：  只能由英文字母、数字组成。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**：  是否开启弹性负载均衡ip组。  **参数范围**：  只能由英文字母、数字组成。
	EnableIpgroup *bool `json:"enable_ipgroup,omitempty"`

	// **参数解释**：  弹性负载均衡ip组类型。  **参数范围**：  只能由英文字母、数字组成。
	Type *string `json:"type,omitempty"`

	// **参数解释**：  ip组信息。  **参数范围**：  只能由英文字母、数字组成。
	Ipgroup        *interface{} `json:"ipgroup,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowIpGroupResponse", string(data)}, " ")
}
