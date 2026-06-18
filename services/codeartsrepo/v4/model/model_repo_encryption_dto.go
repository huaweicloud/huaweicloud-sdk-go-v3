package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoEncryptionDto struct {

	// **参数解释：** 代码仓id。
	RepoId *int32 `json:"repo_id,omitempty"`

	// **参数解释：** 代码仓名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RepoName *string `json:"repo_name,omitempty"`

	// **参数解释：** 代码仓全路径。 **取值范围：** 字符串长度不少于1，不超过1000。
	FullPath *string `json:"full_path,omitempty"`

	// **参数解释：** 项目id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 项目名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 代码仓所有者id。
	OwnerId *int32 `json:"owner_id,omitempty"`

	// **参数解释：** 代码仓所有者iamId。 **取值范围：** 字符串长度不少于1，不超过1000。
	OwnerIamId *string `json:"owner_iam_id,omitempty"`

	// **参数解释：** 代码仓所有者租户id。 **取值范围：** 字符串长度不少于1，不超过1000。
	OwnerTenantName *string `json:"owner_tenant_name,omitempty"`

	// **参数解释：** 代码仓所有者昵称。 **取值范围：** 字符串长度不少于1，不超过1000。
	OwnerNickName *string `json:"owner_nick_name,omitempty"`

	// **参数解释：** 代码仓所有者名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	OwnerName *string `json:"owner_name,omitempty"`
}

func (o RepoEncryptionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoEncryptionDto struct{}"
	}

	return strings.Join([]string{"RepoEncryptionDto", string(data)}, " ")
}
