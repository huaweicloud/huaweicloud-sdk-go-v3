package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessoryUrl struct {

	// 附件id
	AccessoryId *string `json:"accessory_id,omitempty"`

	// 附件链接
	AccessoryUrl *string `json:"accessory_url,omitempty"`
}

func (o AccessoryUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessoryUrl struct{}"
	}

	return strings.Join([]string{"AccessoryUrl", string(data)}, " ")
}
