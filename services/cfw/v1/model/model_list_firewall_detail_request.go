package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFirewallDetailRequest Request Object
type ListFirewallDetailRequest struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 服务类型 0 南北向防火墙 1 东西向防火墙
	ServiceType ListFirewallDetailRequestServiceType `json:"service_type"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，默认情况下，fw_instance_Id为空时，返回账号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 防火墙名称
	Name *string `json:"name,omitempty"`
}

func (o ListFirewallDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallDetailRequest struct{}"
	}

	return strings.Join([]string{"ListFirewallDetailRequest", string(data)}, " ")
}

type ListFirewallDetailRequestServiceType struct {
	value int32
}

type ListFirewallDetailRequestServiceTypeEnum struct {
	E_0 ListFirewallDetailRequestServiceType
	E_1 ListFirewallDetailRequestServiceType
}

func GetListFirewallDetailRequestServiceTypeEnum() ListFirewallDetailRequestServiceTypeEnum {
	return ListFirewallDetailRequestServiceTypeEnum{
		E_0: ListFirewallDetailRequestServiceType{
			value: 0,
		}, E_1: ListFirewallDetailRequestServiceType{
			value: 1,
		},
	}
}

func (c ListFirewallDetailRequestServiceType) Value() int32 {
	return c.value
}

func (c ListFirewallDetailRequestServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFirewallDetailRequestServiceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
