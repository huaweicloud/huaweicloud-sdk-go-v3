package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFoldersAndFilesResponse Response Object
type ListFoldersAndFilesResponse struct {

	// 文件系统所在区域代号。
	Location *string `json:"location,omitempty"`

	// 文件名的绝对路径的前缀。
	Prefix *string `json:"prefix,omitempty"`

	// 文件信息列表。
	Files *[]FileInfo `json:"files,omitempty"`

	// 文件夹信息列表。
	Folders *[]FolderInfo `json:"folders,omitempty"`

	// 本次列举的文件的起始位置，默认为空。
	Marker *string `json:"marker,omitempty"`

	// 下次列举请求的的起始位置，默认为空。
	NextMarker *string `json:"next_marker,omitempty"`

	// 本次列举可以返回的最大文件数量，非本次列举返回的数量，取入参与系统默认值的小值为准。
	MaxKeys        *int32 `json:"max_keys,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFoldersAndFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFoldersAndFilesResponse struct{}"
	}

	return strings.Join([]string{"ListFoldersAndFilesResponse", string(data)}, " ")
}
