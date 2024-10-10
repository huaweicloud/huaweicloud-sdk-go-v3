package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BwListIps struct {

	// 黑白名单ip
	Ip *string `json:"ip,omitempty"`

	// 描述
	Desc *string `json:"desc,omitempty"`
}

func (o BwListIps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BwListIps struct{}"
	}

	return strings.Join([]string{"BwListIps", string(data)}, " ")
}
