package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试用例状态信息，数组长度小于等于50
type UpdateTestCaseResultBean struct {
	// 测试用例唯一标识，列表中不允许存在重复的id，固定长度32位字符

	TestcaseId string `json:"testcase_id"`
	// 注册服务执行id，该值不允许重复，不超过32位字符

	ExecuteId string `json:"execute_id"`
	// 测试用例结果，（0-成功，1-失败，5-执行中，6-停止）

	ResultId string `json:"result_id"`
	// 用例结束执行的时间戳，在执行结束时该字段必传

	EndTime int64 `json:"end_time"`
	// 执行用例持续时长ms，更新状态时改字段必传

	Duration *int64 `json:"duration,omitempty"`
	// 用于记录该次结果执行的备注信息

	Description *string `json:"description,omitempty"`
}

func (o UpdateTestCaseResultBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseResultBean struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseResultBean", string(data)}, " ")
}
