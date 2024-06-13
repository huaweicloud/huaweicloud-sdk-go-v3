package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestItemVo 实际的数据类型：单个对象，集合 或 NULL
type TestItemVo struct {

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

	// 责任人
	Owner *string `json:"owner,omitempty"`

	// frequence值
	Frequence *string `json:"frequence,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

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

	// 是否特性
	IsFeature *string `json:"is_feature,omitempty"`

	// 是否关联特性
	RelateHtsm *string `json:"relate_htsm,omitempty"`

	// aw id
	AwUniqueId *string `json:"aw_unique_id,omitempty"`

	// 脑图id
	TestMindId *string `json:"test_mind_id,omitempty"`

	// 脑图url
	TestMindUrl *string `json:"test_mind_url,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 用例总数
	CaseTotal *int32 `json:"case_total,omitempty"`

	// 执行总数
	ExecdTotal *int32 `json:"execd_total,omitempty"`

	// is_direct_relation
	IsDirectRelation *bool `json:"is_direct_relation,omitempty"`

	// 是否有子特性
	HasChild *bool `json:"has_child,omitempty"`
}

func (o TestItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestItemVo struct{}"
	}

	return strings.Join([]string{"TestItemVo", string(data)}, " ")
}
