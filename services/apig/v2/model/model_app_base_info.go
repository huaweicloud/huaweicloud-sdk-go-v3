package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppBaseInfo struct {

	// 编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述
	Remark *string `json:"remark,omitempty" xml:"remark"`
}

func (o AppBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppBaseInfo struct{}"
	}

	return strings.Join([]string{"AppBaseInfo", string(data)}, " ")
}
