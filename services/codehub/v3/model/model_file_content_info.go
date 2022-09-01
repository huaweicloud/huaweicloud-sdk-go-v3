package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileContentInfo struct {

	// 文件名
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	// 文件大小
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 文件编码
	Encoding *string `json:"encoding,omitempty" xml:"encoding"`

	// sha256编码的文件内容
	ContentSha256 *string `json:"content_sha256,omitempty" xml:"content_sha256"`

	// 分支名
	Ref *string `json:"ref,omitempty" xml:"ref"`

	// blob sha
	BlobId *string `json:"blob_id,omitempty" xml:"blob_id"`

	// 提交对应的SHA id
	CommitId *string `json:"commit_id,omitempty" xml:"commit_id"`

	// 最后一个提交对应的SHA id
	LastCommitId *string `json:"last_commit_id,omitempty" xml:"last_commit_id"`

	// base64编码的文件内容
	Content *string `json:"content,omitempty" xml:"content"`
}

func (o FileContentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileContentInfo struct{}"
	}

	return strings.Join([]string{"FileContentInfo", string(data)}, " ")
}
