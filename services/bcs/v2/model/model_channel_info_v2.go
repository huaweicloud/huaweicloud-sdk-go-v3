package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelInfoV2 struct {

	// 通道名
	Name string `json:"name" xml:"name"`

	// 通道中组织名
	OrgNames []string `json:"org_names" xml:"org_names"`

	// 通道描述
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o ChannelInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelInfoV2 struct{}"
	}

	return strings.Join([]string{"ChannelInfoV2", string(data)}, " ")
}
