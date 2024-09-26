package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionTestCasesChangeStatisticsVo 实际的数据类型：单个对象，集合 或 NULL
type VersionTestCasesChangeStatisticsVo struct {

	// 新增用例数
	AddTestcasesNumber *int32 `json:"add_testcases_number,omitempty"`

	// 复用用例数
	ReuseTestcasesNumber *int32 `json:"reuse_testcases_number,omitempty"`

	// 修改用例数
	ModifyingTestcasesNumber *int32 `json:"modifying_testcases_number,omitempty"`

	// 修改时间时间戳
	UpdateDateTimestamp *int64 `json:"update_date_timestamp,omitempty"`

	// 修改时间
	UpdateDate *sdktime.SdkTime `json:"update_date,omitempty"`
}

func (o VersionTestCasesChangeStatisticsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionTestCasesChangeStatisticsVo struct{}"
	}

	return strings.Join([]string{"VersionTestCasesChangeStatisticsVo", string(data)}, " ")
}
