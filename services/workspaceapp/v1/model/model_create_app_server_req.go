package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppServerReq 创建云服务请求。
type CreateAppServerReq struct {

	// 创建云服务类型，当前仅支持创建云应用：createApps。
	Type string `json:"type"`

	// 服务器组唯一标识。
	ServerGroupId string `json:"server_group_id"`

	// 可用分区。 > - 将服务创建到指定的可用分区，如果不指定则使用系统随机的可用分区。 > - 获取方式详见可用区管理ListAvailabilityZone：\"GET  /v1/{project_id}/availability-zone\"。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 订购数量。
	SubscriptionNum int32 `json:"subscription_num"`

	// 服务对应的网卡信息，当前未使用该字段。
	Nics *[]Nic `json:"nics,omitempty"`

	// OU名称，在对接AD时使用，需提前在AD中创建OU。
	OuName *string `json:"ou_name,omitempty"`

	// 产品ID。 > - 获取方式详见产品套餐管理ListProduct：\"GET /v1/{project_id}/product\"。
	ProductId string `json:"product_id"`

	// 规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 操作系统类型，当前仅支持Windows。
	OsType *string `json:"os_type,omitempty"`

	RootVolume *Volume `json:"root_volume"`

	SchedulerHints *WdhParam `json:"scheduler_hints,omitempty"`

	// 网卡对应的子网ID。
	SubnetId string `json:"subnet_id"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 是否自动升级hda版本。
	UpdateAccessAgent *bool `json:"update_access_agent,omitempty"`

	CreateServerExtendParam *CreateServerExtendParam `json:"create_server_extend_param,omitempty"`
}

func (o CreateAppServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppServerReq struct{}"
	}

	return strings.Join([]string{"CreateAppServerReq", string(data)}, " ")
}
