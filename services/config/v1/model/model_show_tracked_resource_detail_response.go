package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrackedResourceDetailResponse Response Object
type ShowTrackedResourceDetailResponse struct {

	// 资源id
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 云服务名称
	Provider *string `json:"provider,omitempty"`

	// 资源类型
	Type *string `json:"type,omitempty"`

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// Openstack中的项目id
	ProjectId *string `json:"project_id,omitempty"`

	// Openstack中的项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 企业项目id
	EpId *string `json:"ep_id,omitempty"`

	// 企业项目名称
	EpName *string `json:"ep_name,omitempty"`

	// 资源详情校验码
	Checksum *string `json:"checksum,omitempty"`

	// 资源创建时间
	Created *string `json:"created,omitempty"`

	// 资源更新时间
	Updated *string `json:"updated,omitempty"`

	// 资源操作状态
	ProvisioningState *string `json:"provisioning_state,omitempty"`

	// 资源状态，保有中（Normal）/已删除(Deleted)
	State *string `json:"state,omitempty"`

	// 资源Tag
	Tags map[string]string `json:"tags,omitempty"`

	// 资源详细属性
	Properties     map[string]interface{} `json:"properties,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowTrackedResourceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrackedResourceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTrackedResourceDetailResponse", string(data)}, " ")
}
