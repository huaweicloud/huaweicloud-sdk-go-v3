package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCasesChangeStatisticsResponse Response Object
type ShowTestCasesChangeStatisticsResponse struct {

	// 新增用例数
	AddTestcasesNumber *int32 `json:"add_testcases_number,omitempty"`

	// 复用用例数
	ReuseTestcasesNumber *int32 `json:"reuse_testcases_number,omitempty"`

	// 修改用例数
	ModifyingTestcasesNumber *int32 `json:"modifying_testcases_number,omitempty"`

	// 修改时间时间戳
	UpdateDateTimestamp *int64 `json:"update_date_timestamp,omitempty"`

	// 修改时间
	UpdateDate     *sdktime.SdkTime `json:"update_date,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowTestCasesChangeStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCasesChangeStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowTestCasesChangeStatisticsResponse", string(data)}, " ")
}
