package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SelectedUnit 用户在页面中选择的指标单位， 用于后续指标数据回显和计算
type SelectedUnit struct {
}

func (o SelectedUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectedUnit struct{}"
	}

	return strings.Join([]string{"SelectedUnit", string(data)}, " ")
}
