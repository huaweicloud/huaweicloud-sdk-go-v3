package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFoldersAndFilesRequest Request Object
type ListFoldersAndFilesRequest struct {

	// 用户名
	UserName string `json:"user_name"`

	// 个人文件夹ID。(用户拥有多个文件夹时，如果不传个人文件夹id则选择最早创建的文件系统进行文件查询)
	CloudStorageAssignmentId *string `json:"cloud_storage_assignment_id,omitempty"`

	// 查询文件夹路径
	FolderUrl string `json:"folder_url"`

	// 指定一个标识符，从该标识符以后按文件名的字典顺序返回文件列表。
	Marker *string `json:"marker,omitempty"`

	// 列举文件系统的最大数目，返回的文件系统列表将是按照字典顺序的最多前 maxKeys 个，默认取值为1000。
	MaxKeys *int32 `json:"max_keys,omitempty"`
}

func (o ListFoldersAndFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFoldersAndFilesRequest struct{}"
	}

	return strings.Join([]string{"ListFoldersAndFilesRequest", string(data)}, " ")
}
