package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineDirectoryInfo 基线检查策略目录详细信息
type ShowBaselineDirectoryInfo struct {

	// 根据select_type不同的值，得到不同的含义 - 当select_type为check_type，该字段代表check_type（基线的名称） - 当select_type为tag，该字段代表tag（基线检查项的类型）
	Type *string `json:"type,omitempty"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 云安全实践标准   - cis_standard : 通用安全标准
	Standard *string `json:"standard,omitempty"`

	// 基线检查策略三级目录信息
	DataList *[]ShowBaselineDirectoryInfoDataList `json:"data_list,omitempty"`
}

func (o ShowBaselineDirectoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineDirectoryInfo struct{}"
	}

	return strings.Join([]string{"ShowBaselineDirectoryInfo", string(data)}, " ")
}
