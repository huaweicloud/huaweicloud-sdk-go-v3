package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsQueryReq struct {

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略 “tags”字段。 该字段为false或者未提供该参数时，该条件不生效。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含20个key，每个key下面的value最多20个。无tag过滤条件时返回全量数据。
	Tags *[]TmsKeyValues `json:"tags,omitempty"`

	// 企业项目.仅op_service权限可以使用此字段做资源实例过滤条件. 无sys_tags时按照tag接口处理，无tag过滤条件时返回全量数据。
	SysTags *[]TmsKeyValues `json:"sys_tags,omitempty"`

	// 搜索字段,key为要匹配的字段,当前限定为resource_name。value为匹配的值。 根据key的值确认是否需要模糊匹配，如resource_name需要实现前缀搜索。
	Matches *[]TmsMatchesKeyValue `json:"matches,omitempty"`
}

func (o TmsQueryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsQueryReq struct{}"
	}

	return strings.Join([]string{"TmsQueryReq", string(data)}, " ")
}
