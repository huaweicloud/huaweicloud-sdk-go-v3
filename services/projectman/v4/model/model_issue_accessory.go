package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueAccessory struct {

	// 附件id
	AttachmentId *int32 `json:"attachment_id,omitempty" xml:"attachment_id"`

	// 工作鞋ID
	IssueId *int32 `json:"issue_id,omitempty" xml:"issue_id"`

	// 创建者数字ID
	CreatorNumId *int32 `json:"creator_num_id,omitempty" xml:"creator_num_id"`

	// 附件创建时间
	CreatedDate *string `json:"created_date,omitempty" xml:"created_date"`

	// 上传时文件名
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 附件id
	ContainerType *string `json:"container_type,omitempty" xml:"container_type"`

	// 附件名称
	DiskFileName *string `json:"disk_file_name,omitempty" xml:"disk_file_name"`

	// 附件id
	Digest *string `json:"digest,omitempty" xml:"digest"`

	// 附件路径
	DiskDirectory *string `json:"disk_directory,omitempty" xml:"disk_directory"`

	// 创建这用户uuid
	CreatorId *string `json:"creator_id,omitempty" xml:"creator_id"`
}

func (o IssueAccessory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueAccessory struct{}"
	}

	return strings.Join([]string{"IssueAccessory", string(data)}, " ")
}
