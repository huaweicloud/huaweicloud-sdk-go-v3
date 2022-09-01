package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelInfo struct {

	// 通道名
	Name *string `json:"name,omitempty" xml:"name"`

	// 通道中组织名
	OrgNames *[]string `json:"org_names,omitempty" xml:"org_names"`

	// 通道中组织名的哈希值
	OrgNameHash *[]string `json:"org_name_hash,omitempty" xml:"org_name_hash"`

	// key:组织名，value:peer节点数组
	Peers map[string][]string `json:"peers,omitempty" xml:"peers"`
}

func (o ChannelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelInfo struct{}"
	}

	return strings.Join([]string{"ChannelInfo", string(data)}, " ")
}
