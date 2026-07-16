package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferApiKeysRequest Request Object
type ListInferApiKeysRequest struct {

	// **参数解释：** 生效范围，枚举值，可选值为USER、SERVICE。 **约束限制：** 不涉及。 **取值范围：** - USER：表示生效范围为用户级别，可以访问该用户创建的所有在线服务。 - SERVICE：表示生效范围为单个服务，可以访问绑定该api-key的在线服务。 **默认取值：** 不涉及。
	Scope *string `json:"scope,omitempty"`

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	ServiceId *string `json:"service_id,omitempty"`

	// **参数解释：** apikey_name，通过[查询api-keys列表](ListInferApiKeys.xml)获取当前用户拥有的apikey，其中name字段即为apikey_name。 **约束限制：** 不涉及。 **取值范围：** apikey_name。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 服务名，用户在[创建服务](CreateInferService.xml)时自定义。 **约束限制：** 服务名称不能重复。 **取值范围：** 支持1-128个字符，可以包含字母、汉字、数字、连字符和下划线。 **默认取值：** 不涉及。
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释：** apikey_id，在[创建API_KEY](CreateInferApiKey.xml)时即可在返回体中获取，也可通过[查询api-keys列表](ListInferApiKeys.xml)获取当前用户拥有的apikey，其中key_id字段即为apikey_id。 **约束限制：** 不涉及。 **取值范围：** apikey_id只能由英文小写字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及。
	KeyId *string `json:"key_id,omitempty"`

	// **参数解释：** 指定返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表查询的偏移量。 **约束限制：** offset必须是limit的整数倍。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 工作空间ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** 是否查询鉴权范围为USER类型的API KEY，当同时指定service_id时可查询该服务所有可访问的API KEY。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** false。
	WithUserScope *bool `json:"with_user_scope,omitempty"`
}

func (o ListInferApiKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferApiKeysRequest struct{}"
	}

	return strings.Join([]string{"ListInferApiKeysRequest", string(data)}, " ")
}
