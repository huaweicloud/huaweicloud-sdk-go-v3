package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemediationResourceKey 合规规则修正记录查询条件详情。
type RemediationResourceKey struct {

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 云服务名称，为兼容部分合规结果中云服务名称为空的场景，可为空字符串。
	ResourceProvider string `json:"resource_provider"`

	// 资源类型，为兼容部分合规结果中资源类型为空的场景，可为空字符串。
	ResourceType string `json:"resource_type"`
}

func (o RemediationResourceKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationResourceKey struct{}"
	}

	return strings.Join([]string{"RemediationResourceKey", string(data)}, " ")
}
