package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateFlowLogReq
type CreateFlowLogReq struct {

	// 功能说明：流日志名称 取值范围：0-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 功能说明：流日志描述 取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	// 功能说明：流日志所属资源类型 取值范围：支持port、network、vpc 类型。
	ResourceType CreateFlowLogReqResourceType `json:"resource_type"`

	// resource_type对应资源的唯一ID
	ResourceId string `json:"resource_id"`

	// 功能说明：流日志采集类型 取值范围：     1）all：采集指定资源的全部流量。     2）accept：采集指定资源允许传入、传出的流量。     3）reject：采集指定资源拒绝传入、传出的流量。
	TrafficType CreateFlowLogReqTrafficType `json:"traffic_type"`

	// 日志组ID 请在云日志服务中获取，详情请参见《云日志服务用户指南》。
	LogGroupId string `json:"log_group_id"`

	// 日志主题ID 请在云日志服务中获取，详情请参见《云日志服务用户指南》。
	LogTopicId string `json:"log_topic_id"`

	// 功能说明：是否开启日志索引 取值范围：true，false
	IndexEnabled *bool `json:"index_enabled,omitempty"`
}

func (o CreateFlowLogReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowLogReq struct{}"
	}

	return strings.Join([]string{"CreateFlowLogReq", string(data)}, " ")
}

type CreateFlowLogReqResourceType struct {
	value string
}

type CreateFlowLogReqResourceTypeEnum struct {
	PORT    CreateFlowLogReqResourceType
	NETWORK CreateFlowLogReqResourceType
	VPC     CreateFlowLogReqResourceType
}

func GetCreateFlowLogReqResourceTypeEnum() CreateFlowLogReqResourceTypeEnum {
	return CreateFlowLogReqResourceTypeEnum{
		PORT: CreateFlowLogReqResourceType{
			value: "port",
		},
		NETWORK: CreateFlowLogReqResourceType{
			value: "network",
		},
		VPC: CreateFlowLogReqResourceType{
			value: "vpc",
		},
	}
}

func (c CreateFlowLogReqResourceType) Value() string {
	return c.value
}

func (c CreateFlowLogReqResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFlowLogReqResourceType) UnmarshalJSON(b []byte) error {
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

type CreateFlowLogReqTrafficType struct {
	value string
}

type CreateFlowLogReqTrafficTypeEnum struct {
	ALL    CreateFlowLogReqTrafficType
	ACCEPT CreateFlowLogReqTrafficType
	REJECT CreateFlowLogReqTrafficType
}

func GetCreateFlowLogReqTrafficTypeEnum() CreateFlowLogReqTrafficTypeEnum {
	return CreateFlowLogReqTrafficTypeEnum{
		ALL: CreateFlowLogReqTrafficType{
			value: "all",
		},
		ACCEPT: CreateFlowLogReqTrafficType{
			value: "accept",
		},
		REJECT: CreateFlowLogReqTrafficType{
			value: "reject",
		},
	}
}

func (c CreateFlowLogReqTrafficType) Value() string {
	return c.value
}

func (c CreateFlowLogReqTrafficType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFlowLogReqTrafficType) UnmarshalJSON(b []byte) error {
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
