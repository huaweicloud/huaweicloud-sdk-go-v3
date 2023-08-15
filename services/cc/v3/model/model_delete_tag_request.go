package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteTagRequest Request Object
type DeleteTagRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 待删除资源标签的key
	TagKey string `json:"tag_key"`

	// 资源类型: - cloud-connection: 云连接 - bandwidth-package: 带宽包
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
	CLOUD_CONNECTION  DeleteTagRequestResourceType
	BANDWIDTH_PACKAGE DeleteTagRequestResourceType
}

func GetDeleteTagRequestResourceTypeEnum() DeleteTagRequestResourceTypeEnum {
	return DeleteTagRequestResourceTypeEnum{
		CLOUD_CONNECTION: DeleteTagRequestResourceType{
			value: "cloud-connection",
		},
		BANDWIDTH_PACKAGE: DeleteTagRequestResourceType{
			value: "bandwidth-package",
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
