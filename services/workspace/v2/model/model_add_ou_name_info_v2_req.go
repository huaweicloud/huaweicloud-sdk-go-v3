package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddOuNameInfoV2Req struct {

	// 域名称
	Domain string `json:"domain"`

	// OU名称
	OuName string `json:"ou_name"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o AddOuNameInfoV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOuNameInfoV2Req struct{}"
	}

	return strings.Join([]string{"AddOuNameInfoV2Req", string(data)}, " ")
}
