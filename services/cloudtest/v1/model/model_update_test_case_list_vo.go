package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestCaseListVo 实际的数据类型：单个对象，集合 或 NULL
type UpdateTestCaseListVo struct {

	// CTS需要返回资源id
	Id *string `json:"id,omitempty"`

	// CTS需要返回资源name
	Name *string `json:"name,omitempty"`

	// 成功批量更新用例的id列表
	SuccessList *[]string `json:"success_list,omitempty"`

	// 没有批量更新用例的id列表
	FailedList *[]string `json:"failed_list,omitempty"`
}

func (o UpdateTestCaseListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseListVo struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseListVo", string(data)}, " ")
}
