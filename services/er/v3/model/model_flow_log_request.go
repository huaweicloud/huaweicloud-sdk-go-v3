package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlowLogRequest 创建流日志请求体
type FlowLogRequest struct {

	// 流日志名称
	Name string `json:"name"`

	// 流日志描述
	Description *string `json:"description,omitempty"`

	// 流日志采集的资源类型:   - VPC连接   - 虚拟网关连接   - 对等连接
	ResourceType FlowLogRequestResourceType `json:"resource_type"`

	// 要采集的资源ID
	ResourceId string `json:"resource_id"`

	// 日志组ID。请在云日志服务中获取，详情请参见《云日志服务用户指南》。
	LogGroupId string `json:"log_group_id"`

	// 日志主题ID。请在云日志服务中获取，详情请参见《云日志服务用户指南》。
	LogStreamId string `json:"log_stream_id"`

	// 流日志的存储类型: - LTS: 云日志服务器存储
	LogStoreType FlowLogRequestLogStoreType `json:"log_store_type"`
}

func (o FlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowLogRequest struct{}"
	}

	return strings.Join([]string{"FlowLogRequest", string(data)}, " ")
}

type FlowLogRequestResourceType struct {
	value string
}

type FlowLogRequestResourceTypeEnum struct {
	ATTACHMENT FlowLogRequestResourceType
}

func GetFlowLogRequestResourceTypeEnum() FlowLogRequestResourceTypeEnum {
	return FlowLogRequestResourceTypeEnum{
		ATTACHMENT: FlowLogRequestResourceType{
			value: "attachment",
		},
	}
}

func (c FlowLogRequestResourceType) Value() string {
	return c.value
}

func (c FlowLogRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowLogRequestResourceType) UnmarshalJSON(b []byte) error {
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

type FlowLogRequestLogStoreType struct {
	value string
}

type FlowLogRequestLogStoreTypeEnum struct {
	LTS FlowLogRequestLogStoreType
}

func GetFlowLogRequestLogStoreTypeEnum() FlowLogRequestLogStoreTypeEnum {
	return FlowLogRequestLogStoreTypeEnum{
		LTS: FlowLogRequestLogStoreType{
			value: "LTS",
		},
	}
}

func (c FlowLogRequestLogStoreType) Value() string {
	return c.value
}

func (c FlowLogRequestLogStoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowLogRequestLogStoreType) UnmarshalJSON(b []byte) error {
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
