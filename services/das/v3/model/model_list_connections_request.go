package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListConnectionsRequest Request Object
type ListConnectionsRequest struct {

	// 数据库实例地址/实例名称/备注等关键字
	Condition *string `json:"condition,omitempty"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为200。
	Limit *int32 `json:"limit,omitempty"`

	// 数据库来源
	NetworkType *ListConnectionsRequestNetworkType `json:"network_type,omitempty"`

	// 数据库引擎。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 连接类型，NORMAL-创建的连接，SHARE-他人共享给我的连接。
	ConnectionType *ListConnectionsRequestConnectionType `json:"connection_type,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 请求语言类型。
	XLanguage *ListConnectionsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}

type ListConnectionsRequestNetworkType struct {
	value string
}

type ListConnectionsRequestNetworkTypeEnum struct {
	GAUSSDB ListConnectionsRequestNetworkType
	DDS     ListConnectionsRequestNetworkType
	RDS     ListConnectionsRequestNetworkType
	ECS     ListConnectionsRequestNetworkType
	GEMINI  ListConnectionsRequestNetworkType
	DDM     ListConnectionsRequestNetworkType
}

func GetListConnectionsRequestNetworkTypeEnum() ListConnectionsRequestNetworkTypeEnum {
	return ListConnectionsRequestNetworkTypeEnum{
		GAUSSDB: ListConnectionsRequestNetworkType{
			value: "gaussdb",
		},
		DDS: ListConnectionsRequestNetworkType{
			value: "dds",
		},
		RDS: ListConnectionsRequestNetworkType{
			value: "rds",
		},
		ECS: ListConnectionsRequestNetworkType{
			value: "ecs",
		},
		GEMINI: ListConnectionsRequestNetworkType{
			value: "gemini",
		},
		DDM: ListConnectionsRequestNetworkType{
			value: "ddm",
		},
	}
}

func (c ListConnectionsRequestNetworkType) Value() string {
	return c.value
}

func (c ListConnectionsRequestNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConnectionsRequestNetworkType) UnmarshalJSON(b []byte) error {
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

type ListConnectionsRequestConnectionType struct {
	value string
}

type ListConnectionsRequestConnectionTypeEnum struct {
	NORMAL ListConnectionsRequestConnectionType
	SHARE  ListConnectionsRequestConnectionType
}

func GetListConnectionsRequestConnectionTypeEnum() ListConnectionsRequestConnectionTypeEnum {
	return ListConnectionsRequestConnectionTypeEnum{
		NORMAL: ListConnectionsRequestConnectionType{
			value: "NORMAL",
		},
		SHARE: ListConnectionsRequestConnectionType{
			value: "SHARE",
		},
	}
}

func (c ListConnectionsRequestConnectionType) Value() string {
	return c.value
}

func (c ListConnectionsRequestConnectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConnectionsRequestConnectionType) UnmarshalJSON(b []byte) error {
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

type ListConnectionsRequestXLanguage struct {
	value string
}

type ListConnectionsRequestXLanguageEnum struct {
	EN_US ListConnectionsRequestXLanguage
	ZH_CN ListConnectionsRequestXLanguage
}

func GetListConnectionsRequestXLanguageEnum() ListConnectionsRequestXLanguageEnum {
	return ListConnectionsRequestXLanguageEnum{
		EN_US: ListConnectionsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListConnectionsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListConnectionsRequestXLanguage) Value() string {
	return c.value
}

func (c ListConnectionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConnectionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
