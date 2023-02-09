package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type FlowLogResp struct {

	// 流日志资源唯一标识
	Id string `json:"id"`

	// 功能说明：流日志名称 取值范围：0-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）
	Name string `json:"name"`

	// 项目ID
	TenantId string `json:"tenant_id"`

	// 功能说明：流日志描述 取值范围：0-255个字符，不能包含“<”和“>”
	Description string `json:"description"`

	// 功能说明：流日志所属资源类型 取值范围：支持Port、Network、VPC 类型。 约束：当resource_type为Port时，Port资源必须是C3、S3、M3三种虚拟机的网卡。
	ResourceType FlowLogRespResourceType `json:"resource_type"`

	// resource_type对应资源的唯一ID
	ResourceId string `json:"resource_id"`

	// 功能说明：流日志采集类型 取值范围：     1）all：采集指定资源的全部流量。     2）accept：采集指定资源允许传入、传出的流量。     3）reject：采集指定资源拒绝传入、传出的流量。
	TrafficType FlowLogRespTrafficType `json:"traffic_type"`

	// 日志组ID 请在云日志服务中获取，详情请参见《云日志服务用户指南》。
	LogGroupId string `json:"log_group_id"`

	// 日志主题ID 请在云日志服务中获取，详情请参见《云日志服务用户指南》。
	LogTopicId string `json:"log_topic_id"`

	// 功能说明：流日志存储类型 取值范围：     lts：存储类型为云日志服务（LTS）。
	LogStoreType FlowLogRespLogStoreType `json:"log_store_type"`

	// 资源创建时间
	CreatedAt string `json:"created_at"`

	// 最近一次更新资源的时间
	UpdatedAt string `json:"updated_at"`

	// 功能说明：流日志管理 取值范围：若为true，表明开启流日志。若为false，则关闭流日志
	AdminState bool `json:"admin_state"`

	// 功能说明：流日志状态 取值范围：ACTIVE、DOWN、ERROR
	Status FlowLogRespStatus `json:"status"`
}

func (o FlowLogResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowLogResp struct{}"
	}

	return strings.Join([]string{"FlowLogResp", string(data)}, " ")
}

type FlowLogRespResourceType struct {
	value string
}

type FlowLogRespResourceTypeEnum struct {
	PORT    FlowLogRespResourceType
	NETWORK FlowLogRespResourceType
	VPC     FlowLogRespResourceType
}

func GetFlowLogRespResourceTypeEnum() FlowLogRespResourceTypeEnum {
	return FlowLogRespResourceTypeEnum{
		PORT: FlowLogRespResourceType{
			value: "port",
		},
		NETWORK: FlowLogRespResourceType{
			value: "network",
		},
		VPC: FlowLogRespResourceType{
			value: "vpc",
		},
	}
}

func (c FlowLogRespResourceType) Value() string {
	return c.value
}

func (c FlowLogRespResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowLogRespResourceType) UnmarshalJSON(b []byte) error {
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

type FlowLogRespTrafficType struct {
	value string
}

type FlowLogRespTrafficTypeEnum struct {
	ALL    FlowLogRespTrafficType
	ACCEPT FlowLogRespTrafficType
	REJECT FlowLogRespTrafficType
}

func GetFlowLogRespTrafficTypeEnum() FlowLogRespTrafficTypeEnum {
	return FlowLogRespTrafficTypeEnum{
		ALL: FlowLogRespTrafficType{
			value: "all",
		},
		ACCEPT: FlowLogRespTrafficType{
			value: "accept",
		},
		REJECT: FlowLogRespTrafficType{
			value: "reject",
		},
	}
}

func (c FlowLogRespTrafficType) Value() string {
	return c.value
}

func (c FlowLogRespTrafficType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowLogRespTrafficType) UnmarshalJSON(b []byte) error {
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

type FlowLogRespLogStoreType struct {
	value string
}

type FlowLogRespLogStoreTypeEnum struct {
	LTS FlowLogRespLogStoreType
}

func GetFlowLogRespLogStoreTypeEnum() FlowLogRespLogStoreTypeEnum {
	return FlowLogRespLogStoreTypeEnum{
		LTS: FlowLogRespLogStoreType{
			value: "lts",
		},
	}
}

func (c FlowLogRespLogStoreType) Value() string {
	return c.value
}

func (c FlowLogRespLogStoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowLogRespLogStoreType) UnmarshalJSON(b []byte) error {
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

type FlowLogRespStatus struct {
	value string
}

type FlowLogRespStatusEnum struct {
	ACTIVE FlowLogRespStatus
	DOWN   FlowLogRespStatus
	ERROR  FlowLogRespStatus
}

func GetFlowLogRespStatusEnum() FlowLogRespStatusEnum {
	return FlowLogRespStatusEnum{
		ACTIVE: FlowLogRespStatus{
			value: "ACTIVE",
		},
		DOWN: FlowLogRespStatus{
			value: "DOWN",
		},
		ERROR: FlowLogRespStatus{
			value: "ERROR",
		},
	}
}

func (c FlowLogRespStatus) Value() string {
	return c.value
}

func (c FlowLogRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowLogRespStatus) UnmarshalJSON(b []byte) error {
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
