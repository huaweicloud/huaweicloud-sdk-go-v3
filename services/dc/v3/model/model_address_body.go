package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddressBody 专线接入点的位置信息。
type AddressBody struct {

	// 英文名称。
	EnUs *string `json:"en_us,omitempty"`

	// 中文信息。
	ZhCn *string `json:"zh_cn,omitempty"`
}

func (o AddressBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressBody struct{}"
	}

	return strings.Join([]string{"AddressBody", string(data)}, " ")
}
