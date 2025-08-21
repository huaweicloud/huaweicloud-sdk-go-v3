package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CesDimsItemDisplayId 查询ID
type CesDimsItemDisplayId struct {

	// 中文名称
	ZhCn string `json:"zh-cn"`

	// 英文名称
	EnUs string `json:"en-us"`
}

func (o CesDimsItemDisplayId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CesDimsItemDisplayId struct{}"
	}

	return strings.Join([]string{"CesDimsItemDisplayId", string(data)}, " ")
}
