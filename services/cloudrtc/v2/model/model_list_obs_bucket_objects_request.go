package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListObsBucketObjectsRequest Request Object
type ListObsBucketObjectsRequest struct {

	// 要查询的桶名
	Bucket string `json:"bucket"`

	// 对象名前缀，可以理解为文件夹路径
	Prefix *string `json:"prefix,omitempty"`

	// 查询类似，取值“folders”“objects”前者为查询目录，后者为查询对象
	Type ListObsBucketObjectsRequestType `json:"type"`

	// 查询bucket所在的region
	Location ListObsBucketObjectsRequestLocation `json:"location"`
}

func (o ListObsBucketObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListObsBucketObjectsRequest", string(data)}, " ")
}

type ListObsBucketObjectsRequestType struct {
	value string
}

type ListObsBucketObjectsRequestTypeEnum struct {
	FOLDERS ListObsBucketObjectsRequestType
	OBJECTS ListObsBucketObjectsRequestType
}

func GetListObsBucketObjectsRequestTypeEnum() ListObsBucketObjectsRequestTypeEnum {
	return ListObsBucketObjectsRequestTypeEnum{
		FOLDERS: ListObsBucketObjectsRequestType{
			value: "folders",
		},
		OBJECTS: ListObsBucketObjectsRequestType{
			value: "objects",
		},
	}
}

func (c ListObsBucketObjectsRequestType) Value() string {
	return c.value
}

func (c ListObsBucketObjectsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListObsBucketObjectsRequestType) UnmarshalJSON(b []byte) error {
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

type ListObsBucketObjectsRequestLocation struct {
	value string
}

type ListObsBucketObjectsRequestLocationEnum struct {
	CN_NORTH_4 ListObsBucketObjectsRequestLocation
	CN_NORTH_1 ListObsBucketObjectsRequestLocation
	CN_NORTH_5 ListObsBucketObjectsRequestLocation
	CN_NORTH_6 ListObsBucketObjectsRequestLocation
	CN_SOUTH_1 ListObsBucketObjectsRequestLocation
	CN_EAST_2  ListObsBucketObjectsRequestLocation
}

func GetListObsBucketObjectsRequestLocationEnum() ListObsBucketObjectsRequestLocationEnum {
	return ListObsBucketObjectsRequestLocationEnum{
		CN_NORTH_4: ListObsBucketObjectsRequestLocation{
			value: "cn-north-4",
		},
		CN_NORTH_1: ListObsBucketObjectsRequestLocation{
			value: "cn-north-1",
		},
		CN_NORTH_5: ListObsBucketObjectsRequestLocation{
			value: "cn-north-5",
		},
		CN_NORTH_6: ListObsBucketObjectsRequestLocation{
			value: "cn-north-6",
		},
		CN_SOUTH_1: ListObsBucketObjectsRequestLocation{
			value: "cn-south-1",
		},
		CN_EAST_2: ListObsBucketObjectsRequestLocation{
			value: "cn-east-2",
		},
	}
}

func (c ListObsBucketObjectsRequestLocation) Value() string {
	return c.value
}

func (c ListObsBucketObjectsRequestLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListObsBucketObjectsRequestLocation) UnmarshalJSON(b []byte) error {
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
