package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachmentVo struct {

	// 附件类型
	AttachmentType *string `json:"attachment_type,omitempty"`

	// 创建人信息
	CreatedBy *interface{} `json:"created_by,omitempty"`

	// 创建时间
	CreatedDate *string `json:"created_date,omitempty"`

	// 附件描述
	Description *string `json:"description,omitempty"`

	// 存储路径
	DiskDirectory *string `json:"disk_directory,omitempty"`

	// 文件大小
	Filesize *string `json:"filesize,omitempty"`

	// 文件id
	Id *string `json:"id,omitempty"`

	// 更新人信息
	ModifiedBy *interface{} `json:"modified_by,omitempty"`

	// 文件名称
	StoreFilename *string `json:"store_filename,omitempty"`

	// 文件hash名称
	Title *string `json:"title,omitempty"`

	// 工作项id
	WorkitemId *string `json:"workitem_id,omitempty"`
}

func (o AttachmentVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentVo struct{}"
	}

	return strings.Join([]string{"AttachmentVo", string(data)}, " ")
}
