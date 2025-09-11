package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentEntity 附件对象
type AttachmentEntity struct {

	// 附件Id
	Id *string `json:"id,omitempty"`

	// 附件所属工作项Id
	IssueId *string `json:"issue_id,omitempty"`

	// 在服务器存储的相对路径
	DiskDirectory *string `json:"disk_directory,omitempty"`

	// 附件大小，单位为Byte，单个附件最大为200MB
	FileSize *string `json:"file_size,omitempty"`

	// 附件原名称
	StoreFilename *string `json:"store_filename,omitempty"`

	// 附件在数据库中存储的名称。格式为uuid+.扩展名
	Title *string `json:"title,omitempty"`

	ModifiedBy *UserEntity `json:"modified_by,omitempty"`

	CreatedBy *UserEntity `json:"created_by,omitempty"`

	// 附件类型，即文件扩展名
	AttachmentType *string `json:"attachment_type,omitempty"`

	// 附件上传时间的时间戳
	CreatedDate *string `json:"created_date,omitempty"`
}

func (o AttachmentEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentEntity struct{}"
	}

	return strings.Join([]string{"AttachmentEntity", string(data)}, " ")
}
