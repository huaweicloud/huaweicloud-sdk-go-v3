package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestCaseReviewVo 实际的数据类型：单个对象，集合 或 NULL
type TestCaseReviewVo struct {

	// 评审URI
	Uri *string `json:"uri,omitempty"`

	// 用例名称
	TestcaseName *string `json:"testcase_name,omitempty"`

	// 用例编号
	TestcaseNumber *string `json:"testcase_number,omitempty"`

	// 用例阶段
	TestcaseStage *string `json:"testcase_stage,omitempty"`

	// 用例修改时间
	TestcaseLastModified *string `json:"testcase_last_modified,omitempty"`

	// 用例修改时间时间戳
	TestcaseLastModifiedTimestamp *int64 `json:"testcase_last_modified_timestamp,omitempty"`

	// 分支用例URI
	TestcaseUri *string `json:"testcase_uri,omitempty"`

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 版本名称
	VersionName *string `json:"version_name,omitempty"`

	// 评审意见
	Comment *string `json:"comment,omitempty"`

	// 闭环意见
	CloseComment *string `json:"close_comment,omitempty"`

	// 评审人
	Reviewer *string `json:"reviewer,omitempty"`

	// 评审创建时间
	CreationDate *string `json:"creation_date,omitempty"`

	// 评审创建时间时间戳
	CreationDateTimestamp *int64 `json:"creation_date_timestamp,omitempty"`

	// 指定的闭环人列表
	CloseUserIds *[]NameAndIdVo `json:"close_user_ids,omitempty"`

	// 实际闭环人
	ActualClosePerson *string `json:"actual_close_person,omitempty"`

	// 评审状态
	Status *string `json:"status,omitempty"`

	// 评审闭环时间
	CloseDate *string `json:"close_date,omitempty"`

	// 评审闭环时间时间戳
	CloseDateTimestamp *int64 `json:"close_date_timestamp,omitempty"`

	// 期望闭环时间
	ExpectCloseDate *string `json:"expect_close_date,omitempty"`

	// 期望闭环时间时间戳
	ExpectCloseDateTimestamp *int64 `json:"expect_close_date_timestamp,omitempty"`
}

func (o TestCaseReviewVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseReviewVo struct{}"
	}

	return strings.Join([]string{"TestCaseReviewVo", string(data)}, " ")
}
