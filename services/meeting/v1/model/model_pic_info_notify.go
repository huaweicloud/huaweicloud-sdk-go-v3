package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 子画面信息
type PicInfoNotify struct {

	// 多画面中每个画面的编号，编号从1开始
	Index *int32 `json:"index,omitempty" xml:"index"`

	// 每个画面中会话标识，即callNumber。
	Id *[]string `json:"id,omitempty" xml:"id"`

	// 是否为辅流 0： 不是辅流 1： 是辅流
	Share *int32 `json:"share,omitempty" xml:"share"`
}

func (o PicInfoNotify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PicInfoNotify struct{}"
	}

	return strings.Join([]string{"PicInfoNotify", string(data)}, " ")
}
