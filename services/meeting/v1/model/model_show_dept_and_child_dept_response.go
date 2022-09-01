package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDeptAndChildDeptResponse struct {

	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId"`

	// 部门id
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode"`

	// 部门层级
	DeptLevel *int32 `json:"deptLevel,omitempty" xml:"deptLevel"`

	// 部门名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 根部门名路径
	DeptNamePath *string `json:"deptNamePath,omitempty" xml:"deptNamePath"`

	// 是否叶子
	IsLeafNode *bool `json:"isLeafNode,omitempty" xml:"isLeafNode"`

	// 父部门编号
	ParentDeptCode *string `json:"parentDeptCode,omitempty" xml:"parentDeptCode"`

	// 部门编码路径
	DeptCodePath *string `json:"deptCodePath,omitempty" xml:"deptCodePath"`

	// 备注
	Note *string `json:"note,omitempty" xml:"note"`

	// 其他用户对该部门下用户的访问权限
	InPermission *string `json:"inPermission,omitempty" xml:"inPermission"`

	// 该部门下用户访问权限控制
	OutPermission *string `json:"outPermission,omitempty" xml:"outPermission"`

	// 许访问的部门列表,前台回显DTO，id为deptCode,mark为部门名称
	DesignatedOutDeptCodes *[]IdMarkDto `json:"designatedOutDeptCodes,omitempty" xml:"designatedOutDeptCodes"`

	// 子部门列表
	ChildDepts *[]QueryDeptResultDto `json:"childDepts,omitempty" xml:"childDepts"`

	// 部门排序号
	SortLevel      *int32 `json:"sortLevel,omitempty" xml:"sortLevel"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDeptAndChildDeptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeptAndChildDeptResponse struct{}"
	}

	return strings.Join([]string{"ShowDeptAndChildDeptResponse", string(data)}, " ")
}
