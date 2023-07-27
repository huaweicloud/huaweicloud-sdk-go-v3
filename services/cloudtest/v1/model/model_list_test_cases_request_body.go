package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTestCasesRequestBody struct {

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset int32 `json:"offset"`

	// 每页显示的条目数量，最大支持200条
	Limit int32 `json:"limit"`

	// 用例测试执行方式ID
	ExecutionTypeId *int32 `json:"execution_type_id,omitempty"`

	// 分支/测试计划ID，长度11-34位字符（字母和数字）。
	VersionId *string `json:"version_id,omitempty"`
}

func (o ListTestCasesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCasesRequestBody struct{}"
	}

	return strings.Join([]string{"ListTestCasesRequestBody", string(data)}, " ")
}
