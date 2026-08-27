package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListModelGroupResourcesRequest Request Object
type ListModelGroupResourcesRequest struct {

	// 模型组id。
	GroupId string `json:"group_id"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`

	// 资源类型过滤（DESKTOP/DESKTOP_TAG，可选）。
	ResourceType *ListModelGroupResourcesRequestResourceType `json:"resource_type,omitempty"`
}

func (o ListModelGroupResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelGroupResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListModelGroupResourcesRequest", string(data)}, " ")
}

type ListModelGroupResourcesRequestResourceType struct {
	value string
}

type ListModelGroupResourcesRequestResourceTypeEnum struct {
	DESKTOP     ListModelGroupResourcesRequestResourceType
	DESKTOP_TAG ListModelGroupResourcesRequestResourceType
}

func GetListModelGroupResourcesRequestResourceTypeEnum() ListModelGroupResourcesRequestResourceTypeEnum {
	return ListModelGroupResourcesRequestResourceTypeEnum{
		DESKTOP: ListModelGroupResourcesRequestResourceType{
			value: "DESKTOP",
		},
		DESKTOP_TAG: ListModelGroupResourcesRequestResourceType{
			value: "DESKTOP_TAG",
		},
	}
}

func (c ListModelGroupResourcesRequestResourceType) Value() string {
	return c.value
}

func (c ListModelGroupResourcesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListModelGroupResourcesRequestResourceType) UnmarshalJSON(b []byte) error {
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
