package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessoryListAccessories struct {

	// 附件编号
	Id *int64 `json:"id,omitempty"`

	// 附件所属的租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 附件所属的组织
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 附件所属的仓库
	RepoName *string `json:"repo_name,omitempty"`

	// 附件镜像的版本
	SigTag *string `json:"sig_tag,omitempty"`

	// 附件镜像的hash值
	SigDigest *string `json:"sig_digest,omitempty"`

	// 附件关联镜像的hash值
	TargetDigest *string `json:"target_digest,omitempty"`

	// 附件镜像的大小
	Size *int64 `json:"size,omitempty"`

	// 附件的类型
	Type *string `json:"type,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o AccessoryListAccessories) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessoryListAccessories struct{}"
	}

	return strings.Join([]string{"AccessoryListAccessories", string(data)}, " ")
}
