package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BanUrlList 封禁url信息。
type BanUrlList struct {

	// 封禁类型。
	Type *string `json:"type,omitempty"`

	// url信息。
	Url *string `json:"url,omitempty"`

	// 创建时间戳（毫秒）。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间戳（毫秒）。
	ModifyTime *int64 `json:"modify_time,omitempty"`
}

func (o BanUrlList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BanUrlList struct{}"
	}

	return strings.Join([]string{"BanUrlList", string(data)}, " ")
}
