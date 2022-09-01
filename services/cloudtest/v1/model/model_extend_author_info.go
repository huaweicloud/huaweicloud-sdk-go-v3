package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtendAuthorInfo struct {

	// id信息
	Id *string `json:"id,omitempty" xml:"id"`

	// 名称信息
	Name *string `json:"name,omitempty" xml:"name"`

	// 时间信息
	Time *string `json:"time,omitempty" xml:"time"`
}

func (o ExtendAuthorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendAuthorInfo struct{}"
	}

	return strings.Join([]string{"ExtendAuthorInfo", string(data)}, " ")
}
