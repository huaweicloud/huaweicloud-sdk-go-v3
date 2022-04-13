package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagInfo struct {
	// 自定义标签id。

	Id *string `json:"id,omitempty"`
	// 自定义标签名称。

	Name *string `json:"name,omitempty"`
}

func (o TagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagInfo struct{}"
	}

	return strings.Join([]string{"TagInfo", string(data)}, " ")
}
