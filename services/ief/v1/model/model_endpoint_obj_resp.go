package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 端点详情
type EndpointObjResp struct {

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 端点描述，最大长度255，不允许^~#$%&*<>()[]{}'\"\\
	Description string `json:"description"`

	// 端点ID
	Id string `json:"id"`

	// 铂金版实例ID，如果为空则表示是专业版实例。
	IefInstanceId string `json:"ief_instance_id"`

	// 是否共享
	IsShared bool `json:"is_shared"`

	// 端点名称，只允许中文字符、英文字符、数字、下划线、中划线，最大长度64 同一个帐号中创建的端点名唯一
	Name string `json:"name"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 端点的属性，端点需要对外展示的属性，示例： - dis: {\"domain_id\":\"user's domain id\"} - servicebus: {\"service_port\":8080} - apigw: {\"domain_id\":\"user's domain id\"}
	Properties map[string]interface{} `json:"properties"`

	// 端点类型 枚举值： - dis - servicebus - apigw
	Type string `json:"type"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`
}

func (o EndpointObjResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointObjResp struct{}"
	}

	return strings.Join([]string{"EndpointObjResp", string(data)}, " ")
}
