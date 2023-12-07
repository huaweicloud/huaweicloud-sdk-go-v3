package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PercentageDetail Landing Zone设置的进度信息。
type PercentageDetail struct {

	// 进度名称。
	PercentageName *string `json:"percentage_name,omitempty"`

	// 进度状态。
	PercentageStatus *string `json:"percentage_status,omitempty"`
}

func (o PercentageDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PercentageDetail struct{}"
	}

	return strings.Join([]string{"PercentageDetail", string(data)}, " ")
}
