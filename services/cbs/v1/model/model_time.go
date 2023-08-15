package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Time struct {

	// 创建时间
	CreateTime string `json:"create_time"`

	// 更新时间
	UpdateTime string `json:"update_time"`
}

func (o Time) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Time struct{}"
	}

	return strings.Join([]string{"Time", string(data)}, " ")
}
