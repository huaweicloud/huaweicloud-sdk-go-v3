package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteTagRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 待删除资源标签的key
	TagKey string `json:"tag_key"`

	// 资源类型: - cc: 云连接 - bwp: 带宽包
	ResourceType DeleteTagRequestResourceType `json:"resource_type"`
}

func (o DeleteTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagRequest", string(data)}, " ")
}

type DeleteTagRequestResourceType struct {
	value string
}

type DeleteTagRequestResourceTypeEnum struct {
	CC  DeleteTagRequestResourceType
	BWP DeleteTagRequestResourceType
}

func GetDeleteTagRequestResourceTypeEnum() DeleteTagRequestResourceTypeEnum {
	return DeleteTagRequestResourceTypeEnum{
		CC: DeleteTagRequestResourceType{
			value: "cc",
		},
		BWP: DeleteTagRequestResourceType{
			value: "bwp",
		},
	}
}

func (c DeleteTagRequestResourceType) Value() string {
	return c.value
}

func (c DeleteTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
