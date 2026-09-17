package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryNamespaceResp 命名空间详情
type QueryNamespaceResp struct {

	// 命名空间详情
	Name *string `json:"name,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// map类型，key为string,value为string
	Labels map[string]string `json:"labels,omitempty"`

	// 命名空间状态
	Status *string `json:"status,omitempty"`
}

func (o QueryNamespaceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryNamespaceResp struct{}"
	}

	return strings.Join([]string{"QueryNamespaceResp", string(data)}, " ")
}
