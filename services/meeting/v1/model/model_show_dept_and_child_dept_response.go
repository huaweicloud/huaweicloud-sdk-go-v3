package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeptAndChildDeptResponse Response Object
type ShowDeptAndChildDeptResponse struct {

	// 企业ID。
	CorpId *string `json:"corpId,omitempty"`

	// 部门编码。
	DeptCode *string `json:"deptCode,omitempty"`

	// 部门层级。
	DeptLevel *int32 `json:"deptLevel,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`

	// 部门名路径。
	DeptNamePath *string `json:"deptNamePath,omitempty"`

	// 是否叶子节点。
	IsLeafNode *bool `json:"isLeafNode,omitempty"`

	// 父部门编码。
	ParentDeptCode *string `json:"parentDeptCode,omitempty"`

	// 部门编码路径。
	DeptCodePath *string `json:"deptCodePath,omitempty"`

	// 备注。
	Note *string `json:"note,omitempty"`

	// 其他用户对该部门下用户的访问权限。
	InPermission *string `json:"inPermission,omitempty"`

	// 该部门下用户访问权限控制。
	OutPermission *string `json:"outPermission,omitempty"`

	// 允许访问的部门列表。
	DesignatedOutDeptCodes *[]IdMarkDto `json:"designatedOutDeptCodes,omitempty"`

	// 子部门列表。
	ChildDepts *[]QueryDeptResultDto `json:"childDepts,omitempty"`

	// 部门排序号。
	SortLevel      *int32 `json:"sortLevel,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDeptAndChildDeptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeptAndChildDeptResponse struct{}"
	}

	return strings.Join([]string{"ShowDeptAndChildDeptResponse", string(data)}, " ")
}
