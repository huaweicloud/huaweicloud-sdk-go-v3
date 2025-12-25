package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorFilesResponse Response Object
type CreateCollectorFilesResponse struct {

	// UUID
	FieldId *string `json:"field_id,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	Path           *string `json:"path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCollectorFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorFilesResponse struct{}"
	}

	return strings.Join([]string{"CreateCollectorFilesResponse", string(data)}, " ")
}
