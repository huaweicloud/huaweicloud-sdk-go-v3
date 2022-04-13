package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RegionLocales struct {
	// 区域的中文名称。

	ZhCn string `json:"zh-cn"`
	// 区域的英文名称。

	EnUs string `json:"en-us"`
	// 区域的葡萄牙语名称。

	PtBr *string `json:"pt-br,omitempty"`
	// 区域的美国西班牙语名称。

	EsUs *string `json:"es-us,omitempty"`
	// 区域的西班牙语名称。

	EsEs *string `json:"es-es,omitempty"`
}

func (o RegionLocales) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionLocales struct{}"
	}

	return strings.Join([]string{"RegionLocales", string(data)}, " ")
}
