package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddResourceInfo struct {

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 迭代uri
	IteratorUri *string `json:"iterator_uri,omitempty"`

	// 资源类型, 对应serviceType
	Type *int32 `json:"type,omitempty"`

	// 是否选择issues
	IsAllIssues *string `json:"is_all_issues,omitempty"`

	// 是否选择所有用例
	AllImport *bool `json:"all_import,omitempty"`

	// 按照目录引入用例
	FeatureUri *string `json:"feature_uri,omitempty"`

	// 选择的资源列表, 对应sourceCaseUris
	SimpleResourceinfoList *[]SimpleResourceInfo `json:"simple_resourceinfo_list,omitempty"`

	// 反选的资源列表
	InvertSimpleResourceinfoList *[]SimpleResourceInfo `json:"invert_simple_resourceinfo_list,omitempty"`

	// 是否将需求添加到测试计划（不传或者true添加需求到测试计划，false就不添加）
	AddToIterator *bool `json:"add_to_iterator,omitempty"`
}

func (o AddResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResourceInfo struct{}"
	}

	return strings.Join([]string{"AddResourceInfo", string(data)}, " ")
}
