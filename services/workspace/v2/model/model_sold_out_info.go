package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SoldOutInfo 售罄信息。
type SoldOutInfo struct {

	// 售罄产品ID列表。
	Products *[]string `json:"products,omitempty"`
}

func (o SoldOutInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SoldOutInfo struct{}"
	}

	return strings.Join([]string{"SoldOutInfo", string(data)}, " ")
}
