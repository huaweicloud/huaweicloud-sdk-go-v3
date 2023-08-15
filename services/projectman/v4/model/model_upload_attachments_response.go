package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAttachmentsResponse Response Object
type UploadAttachmentsResponse struct {

	// 关联id
	Id *int32 `json:"id,omitempty"`

	// 工作项id
	IssueId *int32 `json:"issue_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 云盘存贮名
	DiskFilename *string `json:"disk_filename,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 文件大小
	Size           *string `json:"size,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"UploadAttachmentsResponse", string(data)}, " ")
}
