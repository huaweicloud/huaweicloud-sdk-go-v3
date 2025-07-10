package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyOuNameInfoV2Req struct {

	// ouid。
	Id *string `json:"id,omitempty"`

	// OU名称。
	OuName string `json:"ou_name"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o ModifyOuNameInfoV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyOuNameInfoV2Req struct{}"
	}

	return strings.Join([]string{"ModifyOuNameInfoV2Req", string(data)}, " ")
}
