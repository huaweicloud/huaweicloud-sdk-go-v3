package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessoryUrl struct {

	// 附件id
	AccessoryId *string `json:"accessory_id,omitempty"`

	// 文件名称
	AccessoryName *string `json:"accessory_name,omitempty"`

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
