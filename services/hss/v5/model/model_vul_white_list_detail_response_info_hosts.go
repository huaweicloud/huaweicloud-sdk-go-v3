package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulWhiteListDetailResponseInfoHosts struct {

	// **参数解释**: 主机id **取值范围**: 字符长度0-128
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度0-128
	HostName *string `json:"host_name,omitempty"`
}

func (o VulWhiteListDetailResponseInfoHosts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulWhiteListDetailResponseInfoHosts struct{}"
	}

	return strings.Join([]string{"VulWhiteListDetailResponseInfoHosts", string(data)}, " ")
}
