package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试用例执行信息，数组长度小于等于50
type TestCaseExecuteBean struct {

	// 测试用例唯一标识，列表中不允许存在重复的id，固定长度32位字符
	TestcaseId string `json:"testcase_id" xml:"testcase_id"`

	// 注册服务执行id，该值不允许重复，不超过32位字符
	ExecuteId string `json:"execute_id" xml:"execute_id"`

	// 测试用例结果，（0-成功，1-失败，5-执行中，6-停止）
	ResultId string `json:"result_id" xml:"result_id"`

	// 用例开始执行的时间戳，在执行开始时该字段必传
	StartTime int64 `json:"start_time" xml:"start_time"`
}

func (o TestCaseExecuteBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseExecuteBean struct{}"
	}

	return strings.Join([]string{"TestCaseExecuteBean", string(data)}, " ")
}
