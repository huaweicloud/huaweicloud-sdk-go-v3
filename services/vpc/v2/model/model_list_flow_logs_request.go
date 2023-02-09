package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListFlowLogsRequest struct {

	// 流日志资源唯一标识
	Id *string `json:"id,omitempty"`

	// 功能说明：流日志名称 取值范围：0-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 项目ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 功能说明：流日志描述 取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	// 功能说明：流日志所属资源类型 取值范围：支持port、network、vpc 3种类型。
	ResourceType *ListFlowLogsRequestResourceType `json:"resource_type,omitempty"`

	// resource_type对应资源的唯一ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 功能说明：流日志采集类型 取值范围：     1）all：采集指定资源的全部流量。     2）accept：采集指定资源允许传入、传出的流量。     3）reject：采集指定资源拒绝传入、传出的流量。
	TrafficType *ListFlowLogsRequestTrafficType `json:"traffic_type,omitempty"`

	// 日志组ID 请在云日志服务中获取，详情请参见《云日志服务用户指南》。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志主题ID 请在云日志服务中获取，详情请参见《云日志服务用户指南》。
	LogTopicId *string `json:"log_topic_id,omitempty"`

	// 功能说明：流日志存储类型 取值范围：     lts：存储类型为云日志服务（LTS）。
	LogStoreType *ListFlowLogsRequestLogStoreType `json:"log_store_type,omitempty"`

	// 功能说明：流日志状态 取值范围：     ACTIVE：开启     DOWN：关闭     ERROR：异常故障
	Status *ListFlowLogsRequestStatus `json:"status,omitempty"`

	// 功能说明：每页返回的个数 取值范围：0 ~ intmax
	Limit *string `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`
}

func (o ListFlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListFlowLogsRequest", string(data)}, " ")
}

type ListFlowLogsRequestResourceType struct {
	value string
}

type ListFlowLogsRequestResourceTypeEnum struct {
	PORT    ListFlowLogsRequestResourceType
	NETWORK ListFlowLogsRequestResourceType
	VPC     ListFlowLogsRequestResourceType
}

func GetListFlowLogsRequestResourceTypeEnum() ListFlowLogsRequestResourceTypeEnum {
	return ListFlowLogsRequestResourceTypeEnum{
		PORT: ListFlowLogsRequestResourceType{
			value: "port",
		},
		NETWORK: ListFlowLogsRequestResourceType{
			value: "network",
		},
		VPC: ListFlowLogsRequestResourceType{
			value: "vpc",
		},
	}
}

func (c ListFlowLogsRequestResourceType) Value() string {
	return c.value
}

func (c ListFlowLogsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowLogsRequestResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListFlowLogsRequestTrafficType struct {
	value string
}

type ListFlowLogsRequestTrafficTypeEnum struct {
	ALL    ListFlowLogsRequestTrafficType
	REJECT ListFlowLogsRequestTrafficType
	ACCEPT ListFlowLogsRequestTrafficType
}

func GetListFlowLogsRequestTrafficTypeEnum() ListFlowLogsRequestTrafficTypeEnum {
	return ListFlowLogsRequestTrafficTypeEnum{
		ALL: ListFlowLogsRequestTrafficType{
			value: "all",
		},
		REJECT: ListFlowLogsRequestTrafficType{
			value: "reject",
		},
		ACCEPT: ListFlowLogsRequestTrafficType{
			value: "accept",
		},
	}
}

func (c ListFlowLogsRequestTrafficType) Value() string {
	return c.value
}

func (c ListFlowLogsRequestTrafficType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowLogsRequestTrafficType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListFlowLogsRequestLogStoreType struct {
	value string
}

type ListFlowLogsRequestLogStoreTypeEnum struct {
	LTS ListFlowLogsRequestLogStoreType
}

func GetListFlowLogsRequestLogStoreTypeEnum() ListFlowLogsRequestLogStoreTypeEnum {
	return ListFlowLogsRequestLogStoreTypeEnum{
		LTS: ListFlowLogsRequestLogStoreType{
			value: "lts",
		},
	}
}

func (c ListFlowLogsRequestLogStoreType) Value() string {
	return c.value
}

func (c ListFlowLogsRequestLogStoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowLogsRequestLogStoreType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListFlowLogsRequestStatus struct {
	value string
}

type ListFlowLogsRequestStatusEnum struct {
	ACTIVE ListFlowLogsRequestStatus
	DOWN   ListFlowLogsRequestStatus
	ERROR  ListFlowLogsRequestStatus
}

func GetListFlowLogsRequestStatusEnum() ListFlowLogsRequestStatusEnum {
	return ListFlowLogsRequestStatusEnum{
		ACTIVE: ListFlowLogsRequestStatus{
			value: "ACTIVE",
		},
		DOWN: ListFlowLogsRequestStatus{
			value: "DOWN",
		},
		ERROR: ListFlowLogsRequestStatus{
			value: "ERROR",
		},
	}
}

func (c ListFlowLogsRequestStatus) Value() string {
	return c.value
}

func (c ListFlowLogsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowLogsRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
