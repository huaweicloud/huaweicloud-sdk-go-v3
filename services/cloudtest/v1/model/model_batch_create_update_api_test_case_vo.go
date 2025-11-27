package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateUpdateApiTestCaseVo 实际的数据类型：单个对象，集合 或 NULL
type BatchCreateUpdateApiTestCaseVo struct {

	// 资源URI
	Uri *string `json:"uri,omitempty"`

	// 资源类型
	Type *string `json:"type,omitempty"`

	// 创建人
	Author *string `json:"author,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 级别
	Rank *int32 `json:"rank,omitempty"`

	// 最后修改人
	LastModifier *string `json:"last_modifier,omitempty"`

	// 最后修改时间
	LastModified *sdktime.SdkTime `json:"last_modified,omitempty"`

	// 修改时间时间戳
	LastModifiedTimestamp *int64 `json:"last_modified_timestamp,omitempty"`

	// 最后变更时间
	LastChangeTime *sdktime.SdkTime `json:"last_change_time,omitempty"`

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 源资源URI
	OriginUri *string `json:"origin_uri,omitempty"`

	// 父资源URI
	ParentUri *string `json:"parent_uri,omitempty"`

	// 父资源路径
	ParentPath *string `json:"parent_path,omitempty"`

	// 创建版本URI
	CreationVersionUri *string `json:"creation_version_uri,omitempty"`

	// 创建时间
	CreationDate *sdktime.SdkTime `json:"creation_date,omitempty"`

	// 创建时间时间戳
	CreationDateTimestamp *int64 `json:"creation_date_timestamp,omitempty"`

	// 创建人名称
	AuthorName *string `json:"author_name,omitempty"`

	// 备注
	Comment *string `json:"comment,omitempty"`

	// 编号
	Number *string `json:"number,omitempty"`

	// 创建成功的用例列表
	SuccessList *[]TestCaseVo `json:"success_list,omitempty"`

	// 创建失败的用例列表
	FailedList *[]TestCaseInfo `json:"failed_list,omitempty"`
}

func (o BatchCreateUpdateApiTestCaseVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateUpdateApiTestCaseVo struct{}"
	}

	return strings.Join([]string{"BatchCreateUpdateApiTestCaseVo", string(data)}, " ")
}
