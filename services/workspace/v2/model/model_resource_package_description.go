package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcePackageDescription struct {

	// 中文描述。
	ZhCn *string `json:"zh_cn,omitempty"`

	// 英文描述。
	EnUs *string `json:"en_us,omitempty"`
}

func (o ResourcePackageDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePackageDescription struct{}"
	}

	return strings.Join([]string{"ResourcePackageDescription", string(data)}, " ")
}
