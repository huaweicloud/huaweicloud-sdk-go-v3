package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApisV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// API编号
	Id *string `json:"id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty"`

	// 请求协议
	ReqProtocol *string `json:"req_protocol,omitempty"`

	// 请求方法
	ReqMethod *string `json:"req_method,omitempty"`

	// 请求路径
	ReqUri *string `json:"req_uri,omitempty"`

	// 授权类型
	AuthType *string `json:"auth_type,omitempty"`

	// 发布的环境编号
	EnvId *string `json:"env_id,omitempty"`

	// API类型
	Type *int32 `json:"type,omitempty"`

	// 指定需要精确匹配查找的参数名称，目前仅支持name、req_uri
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListApisV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisV2Request struct{}"
	}

	return strings.Join([]string{"ListApisV2Request", string(data)}, " ")
}
