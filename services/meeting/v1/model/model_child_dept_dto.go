package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChildDeptDto 子部门详情。
type ChildDeptDto struct {

	// 部门编码。
	DeptCode *string `json:"deptCode,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`

	// 是否为叶子节点（没有子部门的称为叶子节点）。
	IsLeafNodes *bool `json:"isLeafNodes,omitempty"`
}

func (o ChildDeptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChildDeptDto struct{}"
	}

	return strings.Join([]string{"ChildDeptDto", string(data)}, " ")
}
