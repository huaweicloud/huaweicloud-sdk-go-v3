package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoInfoV2 struct {

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 创建者的用户名，在用户是租户的情况下，用户名和租户名相等
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建者的租户名
	DomainName *string `json:"domain_name,omitempty"`

	// 仓库组名(克隆地址中域名后面项目名前的一段 示例：git@repo.alpha.devcloud.inhuawei.com:Demo00228/testword.git  组名：Demo00228 )
	GroupName *string `json:"group_name,omitempty"`

	// 使用 https 克隆仓库时所使用的 url
	HttpsUrl *string `json:"https_url,omitempty"`

	// 用户的 iam user uuid
	IamUserUuid *string `json:"iam_user_uuid,omitempty"`

	// 当前用户是否是仓库的创建者，1：是，0：不是
	IsOwner *int32 `json:"is_owner,omitempty"`

	// 仓库 LFS 容量，单位为M，大于 1024M 则单位为 G
	LfsSize *string `json:"lfs_size,omitempty"`

	// 项目是否被删除
	ProjectIsDeleted *string `json:"project_is_deleted,omitempty"`

	// 项目的uuid
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 仓库主键id
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// 仓库名
	RepositoryName *string `json:"repository_name,omitempty"`

	// 仓库总容量 = 仓库 LFS 容量 + git 库容量，单位为M，大于 1024M 则单位为 G
	RepositorySize *string `json:"repository_size,omitempty"`

	// 仓库uuid(由CreateRepository接口返回)
	RepositoryUuid *string `json:"repository_uuid,omitempty"`

	// 使用 ssh 方式克隆仓库时所使用的 url
	SshUrl *string `json:"ssh_url,omitempty"`

	// 当前用户是否收藏该仓库
	Star *bool `json:"star,omitempty"`

	// 仓库状态， 0：仓库正常创建成功 1：仓库创建中 2：创建失败 3：仓库冻结 4：仓库已经关闭
	Status *int32 `json:"status,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 用户在仓库中的权限:20：只读成员 30：普通成员 40：管理员
	UserRole *int32 `json:"userRole,omitempty"`

	// 是否可见：0私有仓库，20公有仓库
	VisibilityLevel *int32 `json:"visibility_level,omitempty"`

	// web url 路径，访问它将跳转至仓库详情页
	WebUrl *string `json:"web_url,omitempty"`
}

func (o RepoInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoInfoV2 struct{}"
	}

	return strings.Join([]string{"RepoInfoV2", string(data)}, " ")
}
