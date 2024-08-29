package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipResponseBodySpecEgress 出网IP信息。
type ListEipResponseBodySpecEgress struct {

	// 出网IP带宽。
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 是否关闭出网IP。
	Enabled *bool `json:"enabled,omitempty"`

	// 出网IP列表。
	IpList *[]string `json:"ip_list,omitempty"`
}

func (o ListEipResponseBodySpecEgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipResponseBodySpecEgress struct{}"
	}

	return strings.Join([]string{"ListEipResponseBodySpecEgress", string(data)}, " ")
}
