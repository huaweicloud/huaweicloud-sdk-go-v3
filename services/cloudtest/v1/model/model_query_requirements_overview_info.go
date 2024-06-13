package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRequirementsOverviewInfo struct {

	// 筛选迭代ID
	FixedVersionId *string `json:"fixed_version_id,omitempty"`

	// 模块ID
	ModuleId *string `json:"module_id,omitempty"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	PiFilter *PiFilterInfo `json:"pi_filter,omitempty"`
}

func (o QueryRequirementsOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRequirementsOverviewInfo struct{}"
	}

	return strings.Join([]string{"QueryRequirementsOverviewInfo", string(data)}, " ")
}
