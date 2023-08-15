package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateEnterpriseRouter 企业路由器
type CreateEnterpriseRouter struct {

	// 企业路由器实例名称
	Name string `json:"name"`

	// 企业路由器实例描述信息
	Description *string `json:"description,omitempty"`

	// 企业路由器实例的BGP AS号
	Asn int64 `json:"asn"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 计费模式 按需
	ChargeMode *CreateEnterpriseRouterChargeMode `json:"charge_mode,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`

	// 是否开启默认路由表传播，默认false不开启
	EnableDefaultPropagation *bool `json:"enable_default_propagation,omitempty"`

	// 是否开启默认路由表关联，默认false不开启
	EnableDefaultAssociation *bool `json:"enable_default_association,omitempty"`

	// 企业路由器所在的可用区列表
	AvailabilityZoneIds []string `json:"availability_zone_ids"`

	// 是否开启自动接受共享连接创建，默认false不开启
	AutoAcceptSharedAttachments *bool `json:"auto_accept_shared_attachments,omitempty"`
}

func (o CreateEnterpriseRouter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseRouter struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseRouter", string(data)}, " ")
}

type CreateEnterpriseRouterChargeMode struct {
	value string
}

type CreateEnterpriseRouterChargeModeEnum struct {
	POST_PAID CreateEnterpriseRouterChargeMode
	PRE_PAID  CreateEnterpriseRouterChargeMode
}

func GetCreateEnterpriseRouterChargeModeEnum() CreateEnterpriseRouterChargeModeEnum {
	return CreateEnterpriseRouterChargeModeEnum{
		POST_PAID: CreateEnterpriseRouterChargeMode{
			value: "postPaid",
		},
		PRE_PAID: CreateEnterpriseRouterChargeMode{
			value: "prePaid",
		},
	}
}

func (c CreateEnterpriseRouterChargeMode) Value() string {
	return c.value
}

func (c CreateEnterpriseRouterChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEnterpriseRouterChargeMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
