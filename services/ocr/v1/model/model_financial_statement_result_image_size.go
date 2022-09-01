package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FinancialStatementResultImageSize struct {

	// 矫正后图像的高。
	Height *int32 `json:"height,omitempty" xml:"height"`

	// 矫正后图像的宽。
	Width *int32 `json:"width,omitempty" xml:"width"`
}

func (o FinancialStatementResultImageSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FinancialStatementResultImageSize struct{}"
	}

	return strings.Join([]string{"FinancialStatementResultImageSize", string(data)}, " ")
}
