package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentInfo 依赖附件内容
type AttachmentInfo struct {

	// 附件id
	Id *string `json:"id,omitempty"`

	// 附件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件夹.
	FileFolder *string `json:"file_folder,omitempty"`

	// 当前的工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 修改者ID
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改者名称
	ModifierName *string `json:"modifier_name,omitempty"`

	// 是否删除.
	IsDeleted *string `json:"is_deleted,omitempty"`

	// 存储类型.
	StorageType *string `json:"storage_type,omitempty"`

	// 存储url.
	StorageUrl *string `json:"storage_url,omitempty"`
}

func (o AttachmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentInfo struct{}"
	}

	return strings.Join([]string{"AttachmentInfo", string(data)}, " ")
}
