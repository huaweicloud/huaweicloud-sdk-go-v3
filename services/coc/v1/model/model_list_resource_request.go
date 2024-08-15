package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceRequest Request Object
type ListResourceRequest struct {

	// 云服务名称
	Provider string `json:"provider"`

	// 资源类型名称
	Type string `json:"type"`

	// 最大的返回数量
	Limit int32 `json:"limit"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 资源id列表
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// region id
	RegionId *string `json:"region_id,omitempty"`

	// az id
	AzId *string `json:"az_id,omitempty"`

	// ip类型，fixed：内网IP，floating：弹性公网IP
	IpType *string `json:"ip_type,omitempty"`

	// ip
	Ip *string `json:"ip,omitempty"`

	// 资源状态
	Status *string `json:"status,omitempty"`

	// agent状态
	AgentState *string `json:"agent_state,omitempty"`

	// 镜像名称，模糊匹配
	ImageName *string `json:"image_name,omitempty"`

	// 系统类型
	OsType *string `json:"os_type,omitempty"`

	// 标签的值
	Tag *string `json:"tag,omitempty"`

	// 标签的key
	TagKey *string `json:"tag_key,omitempty"`

	// 分组id
	GroupId *string `json:"group_id,omitempty"`

	// 组件id
	ComponentId *string `json:"component_id,omitempty"`

	// 应用id
	ApplicationId *string `json:"application_id,omitempty"`

	// cce集群id
	CceClusterId *string `json:"cce_cluster_id,omitempty"`

	// vpc id
	VpcId *string `json:"vpc_id,omitempty"`

	// 企业项目id
	EpId *string `json:"ep_id,omitempty"`

	// 资源是否已托管
	IsDelegated *bool `json:"is_delegated,omitempty"`
}

func (o ListResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceRequest struct{}"
	}

	return strings.Join([]string{"ListResourceRequest", string(data)}, " ")
}
