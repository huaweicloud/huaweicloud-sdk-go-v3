package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipsRequest Request Object
type ListEipsRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID。此处仅取type为0的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。
	ObjectId string `json:"object_id"`

	// 查询防护EIP列表关键字，可选用弹性公网ID和弹性公网IP
	KeyWord *string `json:"key_word,omitempty"`

	// 防护状态 null-全部 0-开启防护 1-关闭防护
	Status *string `json:"status,omitempty"`

	// 是否同步租户EIP数据 0-不同步 1-同步
	Sync *int32 `json:"sync,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 设备关键字，是eip绑定的资产的名称或id，可通过EIP服务的查询弹性公网IP列表接口获取，通过返回值中的publicips.id（.表示各对象之间层级的区分）获得
	DeviceKey *string `json:"device_key,omitempty"`

	// 地址类型0 ipv4，1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 防火墙关键字，可使用防火墙ID或名称查询，可通过[防火墙ID获取方式](cfw_02_0028.xml)
	FwKeyWord *string `json:"fw_key_word,omitempty"`

	// 弹性公网ip的企业项目id，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，租户未开启企业项目时为0
	EpsId *string `json:"eps_id,omitempty"`

	// 标签列表信息可通过查询EIP服务界面列表标签页签获得
	Tags *string `json:"tags,omitempty"`
}

func (o ListEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipsRequest struct{}"
	}

	return strings.Join([]string{"ListEipsRequest", string(data)}, " ")
}
