package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CesDimsItem CES查询维度信息
type CesDimsItem struct {

	// 查询参数
	Params []string `json:"params"`

	Display *CesDimsItemDisplay `json:"display"`
}

func (o CesDimsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CesDimsItem struct{}"
	}

	return strings.Join([]string{"CesDimsItem", string(data)}, " ")
}
