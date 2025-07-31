package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowBaselineDirectoryInfoDataList struct {

	// 根据select_type不同的值，得到不同的含义 - 当select_type为check_type，该字段代表tag（基线检查项的类型） - 当select_type为tag，该字段代表check_type（基线的名称）
	Name *string `json:"name,omitempty"`

	// 该项是否被选中
	Enable *bool `json:"enable,omitempty"`
}

func (o ShowBaselineDirectoryInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineDirectoryInfoDataList struct{}"
	}

	return strings.Join([]string{"ShowBaselineDirectoryInfoDataList", string(data)}, " ")
}
