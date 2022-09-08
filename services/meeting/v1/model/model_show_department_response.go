package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDepartmentResponse struct {

	// 部门编码
	DeptCode *string `json:"deptCode,omitempty"`

	// 部门名称
	DeptName *string `json:"deptName,omitempty"`

	// 是否为叶子节点（没有子部门的称为叶子节点）
	IsLeafNodes *bool `json:"isLeafNodes,omitempty"`

	// 子部门详情
	ChildDepts     *[]ChildDeptDto `json:"childDepts,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowDepartmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDepartmentResponse struct{}"
	}

	return strings.Join([]string{"ShowDepartmentResponse", string(data)}, " ")
}
