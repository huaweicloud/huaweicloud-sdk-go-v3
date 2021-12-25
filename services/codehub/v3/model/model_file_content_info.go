package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileContentInfo struct {
	// 文件名

	FileName *string `json:"file_name,omitempty"`
	// 文件路径

	FilePath *string `json:"file_path,omitempty"`
	// 文件大小

	Size *int32 `json:"size,omitempty"`
	// 文件编码

	Encoding *string `json:"encoding,omitempty"`
	// sha256编码的文件内容

	ContentSha256 *string `json:"content_sha256,omitempty"`
	// 分支名

	Ref *string `json:"ref,omitempty"`
	// blob sha

	BlobId *string `json:"blob_id,omitempty"`
	// 提交对应的SHA id

	CommitId *string `json:"commit_id,omitempty"`
	// 最后一个提交对应的SHA id

	LastCommitId *string `json:"last_commit_id,omitempty"`
	// base64编码的文件内容

	Content *string `json:"content,omitempty"`
}

func (o FileContentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileContentInfo struct{}"
	}

	return strings.Join([]string{"FileContentInfo", string(data)}, " ")
}
