package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDirectoryTreeResponse struct {
	ParentDirectoryId *string `json:"parent_directory_id,omitempty"`

	DirectoryName *string `json:"directory_name,omitempty"`

	CategoryType *string `json:"category_type,omitempty"`

	DirectoryId *string `json:"directory_id,omitempty"`

	Paging *bool `json:"paging,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Elements *[]TreeNodeElement `json:"elements,omitempty"`

	SubDirectories *[]DirectoryTreeResp `json:"sub_directories,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowDirectoryTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirectoryTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowDirectoryTreeResponse", string(data)}, " ")
}
