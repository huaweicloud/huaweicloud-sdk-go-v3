package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageServerReq 创建镜像实例请求
type CreateImageServerReq struct {

	// 镜像实例名称，名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成，不能有空格 2. 长度范围1~64个字符
	Name string `json:"name"`

	// 镜像实例描述
	Description *string `json:"description,omitempty"`

	RootVolume *Volume `json:"root_volume"`

	ImageRef *ImageRef `json:"image_ref"`

	// 镜像实例所属虚拟私有云唯一标识。
	VpcId string `json:"vpc_id"`

	// 镜像实例网卡对应的子网唯一标识
	SubnetId string `json:"subnet_id"`

	// 镜像实例产品套餐ID
	ProductId string `json:"product_id"`

	// 规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 镜像实例的可用区，空值表示随机选取可用区
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 云应用仓库软件唯一标识请求列表
	AttachApps *[]string `json:"attach_apps,omitempty"`

	// 应用组授权用户， * 限制用户类型：'USER' - 用户
	AuthorizeAccounts []ImageAccountInfo `json:"authorize_accounts"`

	// 组织名称
	OuName *string `json:"ou_name,omitempty"`

	// 是否为vdi单会话模式
	IsVdi *bool `json:"is_vdi,omitempty"`

	SchedulerHints *WdhParam `json:"scheduler_hints,omitempty"`

	ExtraSessionType *ExtraSessionTypeEnum `json:"extra_session_type,omitempty"`

	// 需要付费的会话数，单位/个
	ExtraSessionSize *int32 `json:"extra_session_size,omitempty"`

	RoutePolicy *RoutePolicy `json:"route_policy,omitempty"`

	// 标签信息，最多包含20个key,不允许重复
	Tags *[]TmsTag `json:"tags,omitempty"`

	// **⚠ : 此属性是预留字段，不需要传值，目前镜像产物默认属于default企业项目** 镜像所属的企业项目ID，默认属于default企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参考“[企业中心总览](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/zh-cn_topic_0123692049.html)”。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateImageServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageServerReq struct{}"
	}

	return strings.Join([]string{"CreateImageServerReq", string(data)}, " ")
}
