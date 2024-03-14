package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointWhiteRequestBody 更新终端节点白名单接口请求结构体
type UpdateEndpointWhiteRequestBody struct {

	// 更新或删除用于控制访问终端节点的白名单。此参数可以添加IPv4或CIDR：  - 当取值不为空时，表示将白名单更新为取值所示内容。  - 当取值为空时，表示删除所有白名单。 默认为空列表。
	Whitelist *[]string `json:"whitelist,omitempty"`

	// 是否开启网络ACL隔离。  - true：开启网络ACL隔离  - false：不开启网络ACL隔离 默认值为false。
	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`
}

func (o UpdateEndpointWhiteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointWhiteRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointWhiteRequestBody", string(data)}, " ")
}
