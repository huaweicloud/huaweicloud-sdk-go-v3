package model

import (
	"encoding/json"

	"strings"
)

type ChannelInfoV2 struct {
	// 通道名

	Name string `json:"name"`
	// 通道中组织名

	OrgNames []string `json:"org_names"`
	// 通道描述

	Description *string `json:"description,omitempty"`
}

func (o ChannelInfoV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChannelInfoV2 struct{}"
	}

	return strings.Join([]string{"ChannelInfoV2", string(data)}, " ")
}
