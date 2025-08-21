package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CesDimsItemDisplay 展示内容
type CesDimsItemDisplay struct {
	Id *CesDimsItemDisplayId `json:"id"`
}

func (o CesDimsItemDisplay) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CesDimsItemDisplay struct{}"
	}

	return strings.Join([]string{"CesDimsItemDisplay", string(data)}, " ")
}
