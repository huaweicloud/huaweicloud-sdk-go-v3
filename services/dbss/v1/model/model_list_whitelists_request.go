package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWhitelistsRequest Request Object
type ListWhitelistsRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**： 分页偏移量，从第一条数据偏移offset条数据后开始查询 **约束限制**： 仅支持大于等于0的整数 **取值范围**： 大于等于0 **默认取值**： 默认值为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每页查询记录数。 **约束限制**： 仅支持大于等于10的整数 **取值范围**： 大于等于10小于等于100 **默认取值**： 默认值为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListWhitelistsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhitelistsRequest struct{}"
	}

	return strings.Join([]string{"ListWhitelistsRequest", string(data)}, " ")
}
