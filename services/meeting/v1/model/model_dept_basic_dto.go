package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeptBasicDto 部门基本信息。
type DeptBasicDto struct {

	// 部门编码。
	DeptCode *string `json:"deptCode,omitempty"`

	// 企业id。
	CorpId *string `json:"corpId,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`

	// 部门名称路径。
	DeptNamePath *string `json:"deptNamePath,omitempty"`

	// 父部门编码。
	ParentDeptCode *string `json:"parentDeptCode,omitempty"`
}

func (o DeptBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeptBasicDto struct{}"
	}

	return strings.Join([]string{"DeptBasicDto", string(data)}, " ")
}
