package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageJobType job类型 * `CREATE_SERVER` - 创建镜像实例 * `CREATE_SERVER_IMAGE` - 构建镜像 * `DELETE_SERVER` - 删除镜像实例
type ImageJobType struct {
	value string
}

type ImageJobTypeEnum struct {
	CREATE_SERVER       ImageJobType
	CREATE_SERVER_IMAGE ImageJobType
	DELETE_SERVER       ImageJobType
}

func GetImageJobTypeEnum() ImageJobTypeEnum {
	return ImageJobTypeEnum{
		CREATE_SERVER: ImageJobType{
			value: "CREATE_SERVER",
		},
		CREATE_SERVER_IMAGE: ImageJobType{
			value: "CREATE_SERVER_IMAGE",
		},
		DELETE_SERVER: ImageJobType{
			value: "DELETE_SERVER",
		},
	}
}

func (c ImageJobType) Value() string {
	return c.value
}

func (c ImageJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageJobType) UnmarshalJSON(b []byte) error {
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
