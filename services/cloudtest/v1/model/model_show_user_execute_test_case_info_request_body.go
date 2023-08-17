package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserExecuteTestCaseInfoRequestBody 查询时段内用例的执行情况请求体
type ShowUserExecuteTestCaseInfoRequestBody struct {

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于20000
	Offset int32 `json:"offset"`

	// 每页显示的条目数量,最大支持100条
	Limit int32 `json:"limit"`

	// 用例执行时间段开始
	ExecuteStartTime string `json:"execute_start_time"`

	// 用例执行时间段截止
	ExecuteEndTime string `json:"execute_end_time"`
}

func (o ShowUserExecuteTestCaseInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserExecuteTestCaseInfoRequestBody struct{}"
	}

	return strings.Join([]string{"ShowUserExecuteTestCaseInfoRequestBody", string(data)}, " ")
}
