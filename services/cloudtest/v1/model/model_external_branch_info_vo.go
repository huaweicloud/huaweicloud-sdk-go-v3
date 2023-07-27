package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalBranchInfoVo 分支版本信息
type ExternalBranchInfoVo struct {

	// 分支ID
	Id *string `json:"id,omitempty"`

	// 资源类型
	Type *string `json:"type,omitempty"`

	// 创建人
	Author *string `json:"author,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 最后修改人
	LastModifier *string `json:"last_modifier,omitempty"`

	// 最后修改时间
	LastModified *sdktime.SdkTime `json:"last_modified,omitempty"`

	// 修改时间时间戳
	LastModifiedTimestamp *int64 `json:"last_modified_timestamp,omitempty"`

	// 创建时间
	CreationDate *sdktime.SdkTime `json:"creation_date,omitempty"`

	// 创建时间时间戳
	CreationDateTimestamp *int64 `json:"creation_date_timestamp,omitempty"`

	// 创建人名称
	AuthorName *string `json:"author_name,omitempty"`

	// 是否为基线分支。0表示不是基线分支，1表示是基线分支。
	IsBaseBranch *int32 `json:"is_base_branch,omitempty"`
}

func (o ExternalBranchInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalBranchInfoVo struct{}"
	}

	return strings.Join([]string{"ExternalBranchInfoVo", string(data)}, " ")
}
