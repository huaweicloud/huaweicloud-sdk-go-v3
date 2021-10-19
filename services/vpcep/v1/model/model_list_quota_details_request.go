package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListQuotaDetailsRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 资源类型。 ● endpoint_service：终端节点服 务 ● endpoint：终端节点

	Type *ListQuotaDetailsRequestType `json:"type,omitempty"`
}

func (o ListQuotaDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsRequest", string(data)}, " ")
}

type ListQuotaDetailsRequestType struct {
	value string
}

type ListQuotaDetailsRequestTypeEnum struct {
	ENDPOINT_SERVICE ListQuotaDetailsRequestType
	ENDPOINT         ListQuotaDetailsRequestType
}

func GetListQuotaDetailsRequestTypeEnum() ListQuotaDetailsRequestTypeEnum {
	return ListQuotaDetailsRequestTypeEnum{
		ENDPOINT_SERVICE: ListQuotaDetailsRequestType{
			value: "endpoint_service",
		},
		ENDPOINT: ListQuotaDetailsRequestType{
			value: "endpoint",
		},
	}
}

func (c ListQuotaDetailsRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListQuotaDetailsRequestType) UnmarshalJSON(b []byte) error {
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
