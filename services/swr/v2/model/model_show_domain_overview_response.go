package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainOverviewResponse Response Object
type ShowDomainOverviewResponse struct {

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名称
	DomainName *string `json:"domain_name,omitempty"`

	// 当前租户的命名空间个数
	NamspaceNum *int64 `json:"namspace_num,omitempty"`

	// 当前租户的仓库个数
	RepoNum *int64 `json:"repo_num,omitempty"`

	// 当前租户的镜像个数
	ImageNum *int64 `json:"image_num,omitempty"`

	// 当前租户存储大小
	StoreSpace *float64 `json:"store_space,omitempty"`

	// 当前租户下载流量大小
	DownflowSize   *float64 `json:"downflow_size,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowDomainOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainOverviewResponse", string(data)}, " ")
}
