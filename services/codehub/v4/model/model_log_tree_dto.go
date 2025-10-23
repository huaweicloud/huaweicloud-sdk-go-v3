package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogTreeDto **参数解释：** 文件/文件夹信息
type LogTreeDto struct {

	// **参数解释：** 文件名
	Name *string `json:"name,omitempty"`

	// **参数解释：** 路径
	Path *string `json:"path,omitempty"`

	// **参数解释：** 文件类型
	Type *string `json:"type,omitempty"`

	// **参数解释：** 最近提交信息
	Commit *RepositoryCommitDto `json:"commit,omitempty"`

	// **参数解释：** 文件id
	BlobId *string `json:"blob_id,omitempty"`

	// **参数解释：** 子模块url地址
	SubmoduleUrl *string `json:"submodule_url,omitempty"`

	// **参数解释：** 文件是否受限
	IsLimited *bool `json:"is_limited,omitempty"`

	// **参数解释：** 子模块链接
	SubmoduleLink *string `json:"submodule_link,omitempty"`

	// **参数解释：** 文件md5
	Md5 *string `json:"md5,omitempty"`

	// **参数解释：** 最近提交作者昵称
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 用户名
	UserName *string `json:"user_name,omitempty"`
}

func (o LogTreeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogTreeDto struct{}"
	}

	return strings.Join([]string{"LogTreeDto", string(data)}, " ")
}
