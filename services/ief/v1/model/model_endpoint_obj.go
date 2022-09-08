package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 端点详情
type EndpointObj struct {

	// 端点描述，最大长度255，不允许^~#$%&*<>()[]{}'\"\\
	Description *string `json:"description,omitempty"`

	// 铂金版实例ID，如果为空则表示是专业版实例。
	IefInstanceId *string `json:"ief_instance_id,omitempty"`

	// 端点名称，只允许中文字符、英文字符、数字、下划线、中划线，最大长度64 同一个帐号中创建的端点名唯一
	Name string `json:"name"`

	// 端点的属性，端点需要对外展示的属性，示例： - dis: {\"domain_id\":\"user's domain id\"} - servicebus: {\"service_port\":8080} - apigw: {\"domain_id\":\"user's domain id\"}
	Properties map[string]string `json:"properties"`

	// 端点类型 枚举值： - dis - servicebus - apigw
	Type string `json:"type"`
}

func (o EndpointObj) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointObj struct{}"
	}

	return strings.Join([]string{"EndpointObj", string(data)}, " ")
}
