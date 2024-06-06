package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EnvironmentItem struct {

	// 环境ID。
	Id *string `json:"id,omitempty"`

	// 环境名称。
	Name *string `json:"name,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 环境状态。
	Status *EnvironmentItemStatus `json:"status,omitempty"`

	// 环境附加属性。 - cluster_id：CCE集群ID。 - enterprise_project_id：企业项目ID。 - env_category: 环境种类，当前支持v1、v2，在授权云存储时，v1、v2种类环境有所不同。 - group_name：主环境绑定的SWR组织名称。 - inbound_eip_addr：负载均衡绑定EIP地址。 - namespace：CCE集群命名空间。 - public_elb_id：ELB ID，主环境绑定的负载均衡ID。 - type：环境类型，当前仅支持exclusive类型。 - vpc_id：主环境绑定的VPC ID。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o EnvironmentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentItem struct{}"
	}

	return strings.Join([]string{"EnvironmentItem", string(data)}, " ")
}

type EnvironmentItemStatus struct {
	value string
}

type EnvironmentItemStatusEnum struct {
	CREATING      EnvironmentItemStatus
	FINISH        EnvironmentItemStatus
	DELETING      EnvironmentItemStatus
	FREEZE        EnvironmentItemStatus
	POLICE_FREEZE EnvironmentItemStatus
	DELETE_FAILED EnvironmentItemStatus
}

func GetEnvironmentItemStatusEnum() EnvironmentItemStatusEnum {
	return EnvironmentItemStatusEnum{
		CREATING: EnvironmentItemStatus{
			value: "creating",
		},
		FINISH: EnvironmentItemStatus{
			value: "finish",
		},
		DELETING: EnvironmentItemStatus{
			value: "deleting",
		},
		FREEZE: EnvironmentItemStatus{
			value: "freeze",
		},
		POLICE_FREEZE: EnvironmentItemStatus{
			value: "police_freeze",
		},
		DELETE_FAILED: EnvironmentItemStatus{
			value: "delete_failed",
		},
	}
}

func (c EnvironmentItemStatus) Value() string {
	return c.value
}

func (c EnvironmentItemStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvironmentItemStatus) UnmarshalJSON(b []byte) error {
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
