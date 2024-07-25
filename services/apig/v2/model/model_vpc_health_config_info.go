package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// VpcHealthConfigInfo 健康检查详情。
type VpcHealthConfigInfo struct {

	// 使用以下协议，对VPC中主机执行健康检查： - TCP - HTTP - HTTPS
	Protocol VpcHealthConfigInfoProtocol `json:"protocol"`

	// 健康检查时的目标路径。protocol = http或https时必选
	Path *string `json:"path,omitempty"`

	// 健康检查时的请求方法
	Method *VpcHealthConfigInfoMethod `json:"method,omitempty"`

	// 健康检查的目标端口，缺少或port = 0时为VPC中主机的端口号。  如果此端口存在非0值，则使用此端口进行健康检查。
	Port *int32 `json:"port,omitempty"`

	// 正常阈值。判定VPC通道中主机正常的依据为：连续检查x成功，x为您设置的正常阈值。
	ThresholdNormal int32 `json:"threshold_normal"`

	// 异常阈值。判定VPC通道中主机异常的依据为：连续检查x失败，x为您设置的异常阈值。
	ThresholdAbnormal int32 `json:"threshold_abnormal"`

	// 间隔时间：连续两次检查的间隔时间，单位为秒。必须大于timeout字段取值。
	TimeInterval int32 `json:"time_interval"`

	// 检查目标HTTP响应时，判断成功使用的HTTP响应码。取值范围为100到599之前的任意整数值，支持如下三种格式： - 多个值，如：200,201,202 - 一系列值，如：200-299 - 组合值，如：201,202,210-299 protocol = http时必选
	HttpCode *string `json:"http_code,omitempty"`

	// 是否开启双向认证。如果开启，则使用实例配置中的backend_client_certificate配置项的证书
	EnableClientSsl *bool `json:"enable_client_ssl,omitempty"`

	// 健康检查状态   - 1：可用   - 2：不可用
	Status *VpcHealthConfigInfoStatus `json:"status,omitempty"`

	// 超时时间：检查期间，无响应的时间，单位为秒。必须小于time_interval字段取值。
	Timeout *int32 `json:"timeout,omitempty"`

	// VPC通道的编号
	VpcChannelId *string `json:"vpc_channel_id,omitempty"`

	// 健康检查的编号
	Id *string `json:"id,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o VpcHealthConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcHealthConfigInfo struct{}"
	}

	return strings.Join([]string{"VpcHealthConfigInfo", string(data)}, " ")
}

type VpcHealthConfigInfoProtocol struct {
	value string
}

type VpcHealthConfigInfoProtocolEnum struct {
	TCP   VpcHealthConfigInfoProtocol
	HTTP  VpcHealthConfigInfoProtocol
	HTTPS VpcHealthConfigInfoProtocol
}

func GetVpcHealthConfigInfoProtocolEnum() VpcHealthConfigInfoProtocolEnum {
	return VpcHealthConfigInfoProtocolEnum{
		TCP: VpcHealthConfigInfoProtocol{
			value: "TCP",
		},
		HTTP: VpcHealthConfigInfoProtocol{
			value: "HTTP",
		},
		HTTPS: VpcHealthConfigInfoProtocol{
			value: "HTTPS",
		},
	}
}

func (c VpcHealthConfigInfoProtocol) Value() string {
	return c.value
}

func (c VpcHealthConfigInfoProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcHealthConfigInfoProtocol) UnmarshalJSON(b []byte) error {
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

type VpcHealthConfigInfoMethod struct {
	value string
}

type VpcHealthConfigInfoMethodEnum struct {
	GET  VpcHealthConfigInfoMethod
	HEAD VpcHealthConfigInfoMethod
}

func GetVpcHealthConfigInfoMethodEnum() VpcHealthConfigInfoMethodEnum {
	return VpcHealthConfigInfoMethodEnum{
		GET: VpcHealthConfigInfoMethod{
			value: "GET",
		},
		HEAD: VpcHealthConfigInfoMethod{
			value: "HEAD",
		},
	}
}

func (c VpcHealthConfigInfoMethod) Value() string {
	return c.value
}

func (c VpcHealthConfigInfoMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcHealthConfigInfoMethod) UnmarshalJSON(b []byte) error {
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

type VpcHealthConfigInfoStatus struct {
	value int32
}

type VpcHealthConfigInfoStatusEnum struct {
	E_1 VpcHealthConfigInfoStatus
	E_2 VpcHealthConfigInfoStatus
}

func GetVpcHealthConfigInfoStatusEnum() VpcHealthConfigInfoStatusEnum {
	return VpcHealthConfigInfoStatusEnum{
		E_1: VpcHealthConfigInfoStatus{
			value: 1,
		}, E_2: VpcHealthConfigInfoStatus{
			value: 2,
		},
	}
}

func (c VpcHealthConfigInfoStatus) Value() int32 {
	return c.value
}

func (c VpcHealthConfigInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcHealthConfigInfoStatus) UnmarshalJSON(b []byte) error {
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
