package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipResponseBodySpecIngress 入网IP信息。
type ListEipResponseBodySpecIngress struct {

	// 入网IP带宽
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 是否关闭入网IP。
	Enabled *bool `json:"enabled,omitempty"`

	// 入网IP列表
	IpList *[]string `json:"ip_list,omitempty"`
}

func (o ListEipResponseBodySpecIngress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipResponseBodySpecIngress struct{}"
	}

	return strings.Join([]string{"ListEipResponseBodySpecIngress", string(data)}, " ")
}
