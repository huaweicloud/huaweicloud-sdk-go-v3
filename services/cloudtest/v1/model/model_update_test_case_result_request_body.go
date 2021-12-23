package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新测试用例状态请求体
type UpdateTestCaseResultRequestBody struct {
	// 注册结果返回的服务id

	ServiceId int32 `json:"service_id"`
	// 测试用例状态信息，数组长度小于等于50

	ExecuteList []UpdateTestCaseResultBean `json:"execute_list"`
}

func (o UpdateTestCaseResultRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseResultRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseResultRequestBody", string(data)}, " ")
}
