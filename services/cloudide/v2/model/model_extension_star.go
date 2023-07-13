package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionStar struct {

	// 插件id
	ExtensionId string `json:"extension_id"`

	// 评星内容
	Comment *string `json:"comment,omitempty"`

	// 评星总数
	Stars int32 `json:"stars"`
}

func (o ExtensionStar) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionStar struct{}"
	}

	return strings.Join([]string{"ExtensionStar", string(data)}, " ")
}
