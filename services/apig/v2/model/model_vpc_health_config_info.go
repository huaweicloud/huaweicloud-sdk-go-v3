package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// 健康检查详情，仅VPC通道类型为2时有效。
type VpcHealthConfigInfo struct {
	// 使用以下协议，对VPC中主机执行健康检查。

	Protocol VpcHealthConfigInfoProtocol `json:"protocol"`
	// 健康检查时的目标路径。protocol = http时必选

	Path *string `json:"path,omitempty"`
	// 健康检查时的请求方法

	Method *VpcHealthConfigInfoMethod `json:"method,omitempty"`
	// 健康检查的目标端口，缺省时为VPC中主机的端口号。

	Port *int32 `json:"port,omitempty"`
	// 正常阈值。判定VPC通道中主机正常的依据为：连续检查x成功，x为您设置的正常阈值。

	ThresholdNormal int32 `json:"threshold_normal"`
	// 异常阙值。判定VPC通道中主机异常的依据为：连续检查x失败，x为您设置的异常阈值。

	ThresholdAbnormal int32 `json:"threshold_abnormal"`
	// 间隔时间：连续两次检查的间隔时间，单位为秒。必须大于timeout字段取值。

	TimeInterval int32 `json:"time_interval"`
	// 检查目标HTTP响应时，判断成功使用的HTTP响应码。取值范围为100到599之前的任意整数值，支持如下三种格式： - 多个值，如：200,201,202 - 一系列值，如：200-299 - 组合值，如：201,202,210-299 protocol = http时必选

	HttpCode *string `json:"http_code,omitempty"`
	// 是否开启双向认证。若开启，则使用实例配置中的backend_client_certificate配置项的证书

	EnableClientSsl *bool `json:"enable_client_ssl,omitempty"`
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
	data, err := json.Marshal(o)
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

func (c VpcHealthConfigInfoProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *VpcHealthConfigInfoProtocol) UnmarshalJSON(b []byte) error {
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

func (c VpcHealthConfigInfoMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *VpcHealthConfigInfoMethod) UnmarshalJSON(b []byte) error {
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
