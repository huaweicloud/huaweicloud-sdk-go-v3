package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddTestItemInfo struct {

	// 备注
	Comment *string `json:"comment,omitempty"`

	// 名称
	Name string `json:"name"`

	// 编号
	Number *string `json:"number,omitempty"`

	// 责任人
	Owner *string `json:"owner,omitempty"`

	// 是否为特性,1:特性 2:目录
	IsFeature *string `json:"is_feature,omitempty"`

	// 项目uuid
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 父节点uri
	ParentUri *string `json:"parent_uri,omitempty"`

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 是否忽略名称重复
	IgnoreDuplicateName *bool `json:"ignore_duplicate_name,omitempty"`
}

func (o AddTestItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTestItemInfo struct{}"
	}

	return strings.Join([]string{"AddTestItemInfo", string(data)}, " ")
}
