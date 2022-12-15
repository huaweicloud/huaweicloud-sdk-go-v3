package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DirectoryTreeResp struct {
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
}

func (o DirectoryTreeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectoryTreeResp struct{}"
	}

	return strings.Join([]string{"DirectoryTreeResp", string(data)}, " ")
}
