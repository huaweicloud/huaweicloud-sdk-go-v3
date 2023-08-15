package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiGroupsV2Request Request Object
type ListApiGroupsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// API分组编号
	Id *string `json:"id,omitempty"`

	// API分组名称
	Name *string `json:"name,omitempty"`

	// 集成应用编号
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// 指定需要精确匹配查找的参数名称，多个参数需要支持精确匹配时参数之间使用“,”隔开。  [当前支持name，domain_name。](tag:hws;hws_hk;hcs;fcs;g42;)[目前仅支持API分组名称](tag:Site)
	PreciseSearch *string `json:"precise_search,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`
}

func (o ListApiGroupsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiGroupsV2Request struct{}"
	}

	return strings.Join([]string{"ListApiGroupsV2Request", string(data)}, " ")
}
