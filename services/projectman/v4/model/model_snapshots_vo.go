package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotsVo 快照视图对象
type SnapshotsVo struct {

	// 快照标题。
	Title *string `json:"title,omitempty"`

	// 工作项ID。
	IssueId *string `json:"issue_id,omitempty"`

	// 快照记录工作项。键为工作项类型编码（如 Bug、IR），值为 IssueVO 对象或工作项ID字符串。
	Snapshot2workitem *interface{} `json:"snapshot2workitem,omitempty"`

	CreatedBy *UserVo `json:"created_by,omitempty"`

	ModifiedBy *UserVo `json:"modified_by,omitempty"`

	// 工作项类型。
	Category *string `json:"category,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 工作项父子挂载路径。
	Path *string `json:"path,omitempty"`

	// 区域。
	Region *string `json:"region,omitempty"`

	// 快照ID。
	Id *string `json:"id,omitempty"`

	// 租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 快照创建时间，unix时间戳，单位：毫秒。
	CreatedDate *string `json:"created_date,omitempty"`

	// 快照最后修改时间，unix时间戳，单位：毫秒。
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 项目空间ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 快照类型。
	Type *string `json:"type,omitempty"`

	// 快照基础信息ID。
	SnapBaseInfoId *string `json:"snap_base_info_id,omitempty"`

	// 工作项类型编码。
	IssueCategory *string `json:"issue_category,omitempty"`

	// 父工作项ID。
	ParentId *string `json:"parent_id,omitempty"`

	// 根工作项ID。
	RootId *string `json:"root_id,omitempty"`

	// 父工作项完整路径。
	ParentFullPath *string `json:"parent_full_path,omitempty"`

	// 父工作项路径。
	ParentPath *string `json:"parent_path,omitempty"`

	// 工作项完整路径。
	FullPath *string `json:"full_path,omitempty"`

	// 快照版本号。
	VersionNumber *int32 `json:"version_number,omitempty"`

	// 是否可删除。
	Deletable *bool `json:"deletable,omitempty"`

	// 工作项类型名称。
	CategoryName *string `json:"category_name,omitempty"`
}

func (o SnapshotsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotsVo struct{}"
	}

	return strings.Join([]string{"SnapshotsVo", string(data)}, " ")
}
