package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpeicialConfiguration struct {

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 备忘录
	Note string `json:"note"`
}

func (o SpeicialConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpeicialConfiguration struct{}"
	}

	return strings.Join([]string{"SpeicialConfiguration", string(data)}, " ")
}
