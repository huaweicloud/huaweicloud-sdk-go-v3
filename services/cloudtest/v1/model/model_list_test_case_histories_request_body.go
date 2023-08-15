package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCaseHistoriesRequestBody 查询用例修改历史记录请求体
type ListTestCaseHistoriesRequestBody struct {

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于20000
	Offset int32 `json:"offset"`

	// 每页显示的条目数量,最大支持200条
	Limit int32 `json:"limit"`

	// 版本ID（分支ID或测试计划ID），长度11-34位字符（字母和数字）。
	VersionId string `json:"version_id"`
}

func (o ListTestCaseHistoriesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCaseHistoriesRequestBody struct{}"
	}

	return strings.Join([]string{"ListTestCaseHistoriesRequestBody", string(data)}, " ")
}
