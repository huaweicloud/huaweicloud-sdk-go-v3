package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 发布的部门信息。
type PublishDeptResponseDto struct {

	// 部门编码。
	DeptCode *string `json:"deptCode,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`
}

func (o PublishDeptResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishDeptResponseDto struct{}"
	}

	return strings.Join([]string{"PublishDeptResponseDto", string(data)}, " ")
}
