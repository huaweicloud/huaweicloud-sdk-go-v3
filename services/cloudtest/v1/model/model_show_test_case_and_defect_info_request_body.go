package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCaseAndDefectInfoRequestBody 查询用户用例关联缺陷的统计信息请求体
type ShowTestCaseAndDefectInfoRequestBody struct {

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset int32 `json:"offset"`

	// 每页显示的条目数量,最大支持100条
	Limit int32 `json:"limit"`

	// 分支ID
	BranchId *string `json:"branch_id,omitempty"`

	// 用例创建时间段开始
	CreateTestcaseStartTime string `json:"create_testcase_start_time"`

	// 用例创建时间段截止
	CreateTestcaseEndTime string `json:"create_testcase_end_time"`

	// 缺陷关联时间段开始
	AssociateDefectStartTime *string `json:"associate_defect_start_time,omitempty"`

	// 缺陷关联时间段截止
	AssociateDefectEndTime *string `json:"associate_defect_end_time,omitempty"`
}

func (o ShowTestCaseAndDefectInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseAndDefectInfoRequestBody struct{}"
	}

	return strings.Join([]string{"ShowTestCaseAndDefectInfoRequestBody", string(data)}, " ")
}
