package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Api struct {

	// 目录ID
	CatalogId *string `json:"catalog_id,omitempty"`

	// api 名称
	Name *string `json:"name,omitempty"`

	// api 描述
	Description *string `json:"description,omitempty"`

	// 是否启用访问日志
	LogFlag *bool `json:"log_flag,omitempty"`

	// Api类型
	ApiType *ApiApiType `json:"api_type,omitempty"`

	AuthType *ApiAuthType `json:"auth_type,omitempty"`

	// 发布类型
	PublishType *ApiPublishType `json:"publish_type,omitempty"`

	// api 审核人
	Manager *string `json:"manager,omitempty"`

	// api路径
	Path *string `json:"path,omitempty"`

	// api 协议
	Protocol *ApiProtocol `json:"protocol,omitempty"`

	// 请求类型
	RequestType *ApiRequestType `json:"request_type,omitempty"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`

	// 可见性
	Visibility *ApiVisibility `json:"visibility,omitempty"`

	// API请求参数列表
	RequestParas *[]RequestPara `json:"request_paras,omitempty"`

	DatasourceConfig *DatasourceConfig `json:"datasource_config,omitempty"`

	BackendConfig *BackendConfig `json:"backend_config,omitempty"`
}

func (o Api) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Api struct{}"
	}

	return strings.Join([]string{"Api", string(data)}, " ")
}

type ApiApiType struct {
	value string
}

type ApiApiTypeEnum struct {
	API_TYPE_CREATE   ApiApiType
	API_TYPE_REGISTER ApiApiType
}

func GetApiApiTypeEnum() ApiApiTypeEnum {
	return ApiApiTypeEnum{
		API_TYPE_CREATE: ApiApiType{
			value: "API_TYPE_CREATE",
		},
		API_TYPE_REGISTER: ApiApiType{
			value: "API_TYPE_REGISTER",
		},
	}
}

func (c ApiApiType) Value() string {
	return c.value
}

func (c ApiApiType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiApiType) UnmarshalJSON(b []byte) error {
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

type ApiAuthType struct {
	value string
}

type ApiAuthTypeEnum struct {
	APP  ApiAuthType
	IAM  ApiAuthType
	NONE ApiAuthType
}

func GetApiAuthTypeEnum() ApiAuthTypeEnum {
	return ApiAuthTypeEnum{
		APP: ApiAuthType{
			value: "APP",
		},
		IAM: ApiAuthType{
			value: "IAM",
		},
		NONE: ApiAuthType{
			value: "NONE",
		},
	}
}

func (c ApiAuthType) Value() string {
	return c.value
}

func (c ApiAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthType) UnmarshalJSON(b []byte) error {
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

type ApiPublishType struct {
	value string
}

type ApiPublishTypeEnum struct {
	PUBLISH_TYPE_PUBLIC  ApiPublishType
	PUBLISH_TYPE_PRIVATE ApiPublishType
}

func GetApiPublishTypeEnum() ApiPublishTypeEnum {
	return ApiPublishTypeEnum{
		PUBLISH_TYPE_PUBLIC: ApiPublishType{
			value: "PUBLISH_TYPE_PUBLIC",
		},
		PUBLISH_TYPE_PRIVATE: ApiPublishType{
			value: "PUBLISH_TYPE_PRIVATE",
		},
	}
}

func (c ApiPublishType) Value() string {
	return c.value
}

func (c ApiPublishType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPublishType) UnmarshalJSON(b []byte) error {
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

type ApiProtocol struct {
	value string
}

type ApiProtocolEnum struct {
	PROTOCOL_TYPE_HTTP  ApiProtocol
	PROTOCOL_TYPE_HTTPS ApiProtocol
}

func GetApiProtocolEnum() ApiProtocolEnum {
	return ApiProtocolEnum{
		PROTOCOL_TYPE_HTTP: ApiProtocol{
			value: "PROTOCOL_TYPE_HTTP",
		},
		PROTOCOL_TYPE_HTTPS: ApiProtocol{
			value: "PROTOCOL_TYPE_HTTPS",
		},
	}
}

func (c ApiProtocol) Value() string {
	return c.value
}

func (c ApiProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiProtocol) UnmarshalJSON(b []byte) error {
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

type ApiRequestType struct {
	value string
}

type ApiRequestTypeEnum struct {
	REQUEST_TYPE_POST ApiRequestType
	REQUEST_TYPE_GET  ApiRequestType
}

func GetApiRequestTypeEnum() ApiRequestTypeEnum {
	return ApiRequestTypeEnum{
		REQUEST_TYPE_POST: ApiRequestType{
			value: "REQUEST_TYPE_POST",
		},
		REQUEST_TYPE_GET: ApiRequestType{
			value: "REQUEST_TYPE_GET",
		},
	}
}

func (c ApiRequestType) Value() string {
	return c.value
}

func (c ApiRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRequestType) UnmarshalJSON(b []byte) error {
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

type ApiVisibility struct {
	value string
}

type ApiVisibilityEnum struct {
	WORKSPACE ApiVisibility
	PROJECT   ApiVisibility
	DOMAIN    ApiVisibility
}

func GetApiVisibilityEnum() ApiVisibilityEnum {
	return ApiVisibilityEnum{
		WORKSPACE: ApiVisibility{
			value: "WORKSPACE",
		},
		PROJECT: ApiVisibility{
			value: "PROJECT",
		},
		DOMAIN: ApiVisibility{
			value: "DOMAIN",
		},
	}
}

func (c ApiVisibility) Value() string {
	return c.value
}

func (c ApiVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiVisibility) UnmarshalJSON(b []byte) error {
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
