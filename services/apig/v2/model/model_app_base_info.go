package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppBaseInfo struct {
	// 编号

	Id *string `json:"id,omitempty"`
	// 名称

	Name *string `json:"name,omitempty"`
	// 描述

	Remark *string `json:"remark,omitempty"`
}

func (o AppBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppBaseInfo struct{}"
	}

	return strings.Join([]string{"AppBaseInfo", string(data)}, " ")
}
