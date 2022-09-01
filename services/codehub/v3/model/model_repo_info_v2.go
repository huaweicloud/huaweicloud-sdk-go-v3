package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoInfoV2 struct {

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 创建者的用户名，在用户是租户的情况下，用户名和租户名相等
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name"`

	// 创建者的租户名
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 仓库组名
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 使用 https 克隆仓库时所使用的 url
	HttpsUrl *string `json:"https_url,omitempty" xml:"https_url"`

	// 用户的 iam user uuid
	IamUserUuid *string `json:"iam_user_uuid,omitempty" xml:"iam_user_uuid"`

	// 当前用户是否是仓库的创建者，1：是，0：不是
	IsOwner *int32 `json:"is_owner,omitempty" xml:"is_owner"`

	// 仓库 LFS 容量，单位为M，大于 1024M 则单位为 G
	LfsSize *string `json:"lfs_size,omitempty" xml:"lfs_size"`

	// 项目是否被删除
	ProjectIsDeleted *string `json:"project_is_deleted,omitempty" xml:"project_is_deleted"`

	// 项目的uuid
	ProjectUuid *string `json:"project_uuid,omitempty" xml:"project_uuid"`

	// 仓库主键id
	RepositoryId *int32 `json:"repository_id,omitempty" xml:"repository_id"`

	// 仓库名
	RepositoryName *string `json:"repository_name,omitempty" xml:"repository_name"`

	// 仓库总容量 = 仓库 LFS 容量 + git 库容量，单位为M，大于 1024M 则单位为 G
	RepositorySize *string `json:"repository_size,omitempty" xml:"repository_size"`

	// 仓库uuid
	RepositoryUuid *string `json:"repository_uuid,omitempty" xml:"repository_uuid"`

	// 使用 ssh 方式克隆仓库时所使用的 url
	SshUrl *string `json:"ssh_url,omitempty" xml:"ssh_url"`

	// 当前用户是否收藏该仓库
	Star *bool `json:"star,omitempty" xml:"star"`

	// 仓库状态， 0：仓库正常创建成功 1：仓库创建中 2：创建失败 3：仓库冻结 4：仓库已经关闭
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 用户在仓库中的权限:20：只读成员 30：普通成员 40：管理员
	UserRole *int32 `json:"userRole,omitempty" xml:"userRole"`

	// 是否可见：0私有仓库，20公有仓库
	VisibilityLevel *int32 `json:"visibility_level,omitempty" xml:"visibility_level"`

	// web url 路径，访问它将跳转至仓库详情页
	WebUrl *string `json:"web_url,omitempty" xml:"web_url"`
}

func (o RepoInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoInfoV2 struct{}"
	}

	return strings.Join([]string{"RepoInfoV2", string(data)}, " ")
}
