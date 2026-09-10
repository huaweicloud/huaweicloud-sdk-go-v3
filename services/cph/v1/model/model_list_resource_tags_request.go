package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceTagsRequest Request Object
type ListResourceTagsRequest struct {

	// 资源类型。  - cph-server，云手机服务器 [- elastic-cloud-phone, 弹性云手机实例](tag:hws_test)
	ResourceType ListResourceTagsRequestResourceType `json:"resource_type"`

	// 资源ID。
	ResourceId string `json:"resource_id"`
}

func (o ListResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceTagsRequest", string(data)}, " ")
}

type ListResourceTagsRequestResourceType struct {
	value string
}

type ListResourceTagsRequestResourceTypeEnum struct {
	CPH_SERVER          ListResourceTagsRequestResourceType
	ELASTIC_CLOUD_PHONE ListResourceTagsRequestResourceType
}

func GetListResourceTagsRequestResourceTypeEnum() ListResourceTagsRequestResourceTypeEnum {
	return ListResourceTagsRequestResourceTypeEnum{
		CPH_SERVER: ListResourceTagsRequestResourceType{
			value: "cph-server",
		},
		ELASTIC_CLOUD_PHONE: ListResourceTagsRequestResourceType{
			value: "elastic-cloud-phone",
		},
	}
}

func (c ListResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
