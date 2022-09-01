package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDeptResultDto struct {

	// 部门编码，企业内唯一
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode"`

	// 部门层级
	DeptLevel *int32 `json:"deptLevel,omitempty" xml:"deptLevel"`

	// 部门名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 部门名路径
	DeptNamePath *string `json:"deptNamePath,omitempty" xml:"deptNamePath"`

	// 是否叶子
	IsLeafNode *bool `json:"isLeafNode,omitempty" xml:"isLeafNode"`

	// 父部门编号
	ParentDeptCode *string `json:"parentDeptCode,omitempty" xml:"parentDeptCode"`

	// 部门编码路径
	DeptCodePath *string `json:"deptCodePath,omitempty" xml:"deptCodePath"`

	// 备注
	Note *string `json:"note,omitempty" xml:"note"`

	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId"`

	// 其他用户对该部门下用户的访问权限
	InPermission *string `json:"inPermission,omitempty" xml:"inPermission"`

	// 该部门下用户访问权限控制
	OutPermission *string `json:"outPermission,omitempty" xml:"outPermission"`

	// 允许访问的部门列表，id为部门编码。
	DesignatedOutDeptCodes *[]IdMarkDto `json:"designatedOutDeptCodes,omitempty" xml:"designatedOutDeptCodes"`

	// 部门排序号
	SortLevel *int32 `json:"sortLevel,omitempty" xml:"sortLevel"`
}

func (o QueryDeptResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDeptResultDto struct{}"
	}

	return strings.Join([]string{"QueryDeptResultDto", string(data)}, " ")
}
