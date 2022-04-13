package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListL7rulesRequest struct {
	// 分页查询中每页的转发规则个数

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询的起始的资源id，表示上一页最后一条查询记录的转发规则的id。不指定时表示查询第一页。

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 转发策略id

	L7policyId string `json:"l7policy_id"`
	// 转发规则ID。

	Id *string `json:"id,omitempty"`
	// 转发规则的管理状态；取值范围： true/false。该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发规则的匹配类型。取值范围：HOST_NAME：匹配请求中的域名；PATH：匹配请求中的路径；同一个转发策略下转发规则的type不能重复。

	Type *string `json:"type,omitempty"`
	// 转发匹配方式： type为HOST_NAME时，取值范围：EQUAL_TO：精确匹配； type为PATH时，取值范围：REGEX：正则匹配；STARTS_WITH：前缀匹配；EQUAL_TO：精确匹配。

	CompareType *string `json:"compare_type,omitempty"`
	// 是否反向匹配；取值范围：true/false。默认值：false；该字段为预留字段，暂未启用。

	Invert *bool `json:"invert,omitempty"`
	// 匹配内容的键值。默认为null。该字段为预留字段，暂未启用。

	Key *string `json:"key,omitempty"`
	// 匹配内容的值。 当type为HOST_NAME时，取值范围：String (100)，字符串只能包含英文字母、数字、“-”或“.”，且必须以字母或数字开头。 当type为PATH时，取值范围：String (128)。当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:| /()[]{}，且必须以\"/\"开头。

	Value *string `json:"value,omitempty"`
	// 转发规则的配置状态，可以为ACTIVE、PENDING_CREATE 或者ERROR。默认值：ACTIVE；该字段为预留字段，暂未启用。

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
}

func (o ListL7rulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7rulesRequest struct{}"
	}

	return strings.Join([]string{"ListL7rulesRequest", string(data)}, " ")
}
