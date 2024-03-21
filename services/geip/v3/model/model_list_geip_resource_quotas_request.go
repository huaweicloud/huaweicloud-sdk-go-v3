package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGeipResourceQuotasRequest Request Object
type ListGeipResourceQuotasRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]string `json:"fields,omitempty"`

	Type *[]ListGeipResourceQuotasRequestType `json:"type,omitempty"`
}

func (o ListGeipResourceQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipResourceQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListGeipResourceQuotasRequest", string(data)}, " ")
}

type ListGeipResourceQuotasRequestType struct {
	value string
}

type ListGeipResourceQuotasRequestTypeEnum struct {
	GEIP                  ListGeipResourceQuotasRequestType
	GEIP_SEGMENT          ListGeipResourceQuotasRequestType
	INTERNET_BANDWIDTH_IP ListGeipResourceQuotasRequestType
	INTERNET_BANDWIDTH    ListGeipResourceQuotasRequestType
}

func GetListGeipResourceQuotasRequestTypeEnum() ListGeipResourceQuotasRequestTypeEnum {
	return ListGeipResourceQuotasRequestTypeEnum{
		GEIP: ListGeipResourceQuotasRequestType{
			value: "geip",
		},
		GEIP_SEGMENT: ListGeipResourceQuotasRequestType{
			value: "geip_segment",
		},
		INTERNET_BANDWIDTH_IP: ListGeipResourceQuotasRequestType{
			value: "internetBandwidthIP",
		},
		INTERNET_BANDWIDTH: ListGeipResourceQuotasRequestType{
			value: "internetBandwidth",
		},
	}
}

func (c ListGeipResourceQuotasRequestType) Value() string {
	return c.value
}

func (c ListGeipResourceQuotasRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipResourceQuotasRequestType) UnmarshalJSON(b []byte) error {
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
