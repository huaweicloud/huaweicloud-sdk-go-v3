package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFileResponse Response Object
type ShowFileResponse struct {

	// 文件名称
	Name *string `json:"name,omitempty"`

	// 文件路径
	Path *string `json:"path,omitempty"`

	// 文件大小
	Size *int64 `json:"size,omitempty"`

	// 文件编码方式
	Encoding *string `json:"encoding,omitempty"`

	// 分支名称、tag名称或者commitid
	Ref *string `json:"ref,omitempty"`

	// 文件标识
	BlobId *string `json:"blob_id,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	Commit *RepositoryCommitDto `json:"commit,omitempty"`

	// 文件内容
	Content *string `json:"content,omitempty"`

	// **参数解释：** 文件是否受限。 **取值范围：** - false, 不受限。 - true, 受限。
	IsLimited *bool `json:"is_limited,omitempty"`

	// sha256值
	ContentSha256 *string `json:"content_sha256,omitempty"`

	// 最新提交commitid
	LastCommitId *string `json:"last_commit_id,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileResponse struct{}"
	}

	return strings.Join([]string{"ShowFileResponse", string(data)}, " ")
}
