package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeCheck任务语言参数
type ParamTypeLimits struct {
	// 是否废弃

	Disable string `json:"disable"`
	// 语言名字

	Name string `json:"name"`
	// 语言展示名字

	Displayname string `json:"displayname"`
	// 规则集ID

	Id string `json:"id"`
	// 扫描语言

	Language string `json:"language"`
}

func (o ParamTypeLimits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamTypeLimits struct{}"
	}

	return strings.Join([]string{"ParamTypeLimits", string(data)}, " ")
}
