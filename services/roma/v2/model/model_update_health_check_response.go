package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// Response Object
type UpdateHealthCheckResponse struct {
	// 使用以下协议，对VPC中主机执行健康检查： - TCP - HTTP - HTTPS

	Protocol UpdateHealthCheckResponseProtocol `json:"protocol"`
	// 健康检查时的目标路径。protocol = http或https时必选

	Path *string `json:"path,omitempty"`
	// 健康检查时的请求方法

	Method *UpdateHealthCheckResponseMethod `json:"method,omitempty"`
	// 健康检查的目标端口，缺少或port = 0时为VPC中主机的端口号。  若此端口存在非0值，则使用此端口进行健康检查。

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
	// 健康检查状态   - 1：可用   - 2：不可用

	Status *UpdateHealthCheckResponseStatus `json:"status,omitempty"`
	// 超时时间：检查期间，无响应的时间，单位为秒。必须小于time_interval字段取值。

	Timeout *int32 `json:"timeout,omitempty"`
	// VPC通道的编号

	VpcChannelId *string `json:"vpc_channel_id,omitempty"`
	// 健康检查的编号

	Id *string `json:"id,omitempty"`
	// 创建时间

	CreateTime     *sdktime.SdkTime `json:"create_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateHealthCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthCheckResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthCheckResponse", string(data)}, " ")
}

type UpdateHealthCheckResponseProtocol struct {
	value string
}

type UpdateHealthCheckResponseProtocolEnum struct {
	TCP   UpdateHealthCheckResponseProtocol
	HTTP  UpdateHealthCheckResponseProtocol
	HTTPS UpdateHealthCheckResponseProtocol
}

func GetUpdateHealthCheckResponseProtocolEnum() UpdateHealthCheckResponseProtocolEnum {
	return UpdateHealthCheckResponseProtocolEnum{
		TCP: UpdateHealthCheckResponseProtocol{
			value: "TCP",
		},
		HTTP: UpdateHealthCheckResponseProtocol{
			value: "HTTP",
		},
		HTTPS: UpdateHealthCheckResponseProtocol{
			value: "HTTPS",
		},
	}
}

func (c UpdateHealthCheckResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHealthCheckResponseProtocol) UnmarshalJSON(b []byte) error {
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

type UpdateHealthCheckResponseMethod struct {
	value string
}

type UpdateHealthCheckResponseMethodEnum struct {
	GET  UpdateHealthCheckResponseMethod
	HEAD UpdateHealthCheckResponseMethod
}

func GetUpdateHealthCheckResponseMethodEnum() UpdateHealthCheckResponseMethodEnum {
	return UpdateHealthCheckResponseMethodEnum{
		GET: UpdateHealthCheckResponseMethod{
			value: "GET",
		},
		HEAD: UpdateHealthCheckResponseMethod{
			value: "HEAD",
		},
	}
}

func (c UpdateHealthCheckResponseMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHealthCheckResponseMethod) UnmarshalJSON(b []byte) error {
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

type UpdateHealthCheckResponseStatus struct {
	value int32
}

type UpdateHealthCheckResponseStatusEnum struct {
	E_1 UpdateHealthCheckResponseStatus
	E_2 UpdateHealthCheckResponseStatus
}

func GetUpdateHealthCheckResponseStatusEnum() UpdateHealthCheckResponseStatusEnum {
	return UpdateHealthCheckResponseStatusEnum{
		E_1: UpdateHealthCheckResponseStatus{
			value: 1,
		}, E_2: UpdateHealthCheckResponseStatus{
			value: 2,
		},
	}
}

func (c UpdateHealthCheckResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHealthCheckResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
