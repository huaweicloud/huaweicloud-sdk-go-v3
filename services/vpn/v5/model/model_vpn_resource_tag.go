package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnResourceTag struct {

	// 标签的key
	Key string `json:"key"`

	// 标签的value
	Value string `json:"value"`
}

func (o VpnResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnResourceTag struct{}"
	}

	return strings.Join([]string{"VpnResourceTag", string(data)}, " ")
}
