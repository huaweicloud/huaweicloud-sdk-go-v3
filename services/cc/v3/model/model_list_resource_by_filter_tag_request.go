package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceByFilterTagRequest Request Object
type ListResourceByFilterTagRequest struct {

	// 资源类型: - cloud-connection: 云连接 - bandwidth-package: 带宽包
	ResourceType ListResourceByFilterTagRequestResourceType `json:"resource_type"`

	Body *ListResourceByFilterTagRequestBody `json:"body,omitempty"`
}

func (o ListResourceByFilterTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceByFilterTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourceByFilterTagRequest", string(data)}, " ")
}

type ListResourceByFilterTagRequestResourceType struct {
	value string
}

type ListResourceByFilterTagRequestResourceTypeEnum struct {
	CLOUD_CONNECTION  ListResourceByFilterTagRequestResourceType
	BANDWIDTH_PACKAGE ListResourceByFilterTagRequestResourceType
}

func GetListResourceByFilterTagRequestResourceTypeEnum() ListResourceByFilterTagRequestResourceTypeEnum {
	return ListResourceByFilterTagRequestResourceTypeEnum{
		CLOUD_CONNECTION: ListResourceByFilterTagRequestResourceType{
			value: "cloud-connection",
		},
		BANDWIDTH_PACKAGE: ListResourceByFilterTagRequestResourceType{
			value: "bandwidth-package",
		},
	}
}

func (c ListResourceByFilterTagRequestResourceType) Value() string {
	return c.value
}

func (c ListResourceByFilterTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceByFilterTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
