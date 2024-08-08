package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageServerStatus 镜像实例的状态: * 'CREATING' - 实例创建中 * 'ACTIVE' - 实例正常运行 * 'BUILDING' - 镜像创建中 * 'BUILT' - 镜像任务结束 * 'ERROR' - 实例处于异常 * 'DELETING' - 实例删除中 * 'NULL' - 未设置
type ImageServerStatus struct {
	value string
}

type ImageServerStatusEnum struct {
	CREATING ImageServerStatus
	ACTIVE   ImageServerStatus
	BUILDING ImageServerStatus
	BUILT    ImageServerStatus
	ERROR    ImageServerStatus
	DELETING ImageServerStatus
	NULL     ImageServerStatus
}

func GetImageServerStatusEnum() ImageServerStatusEnum {
	return ImageServerStatusEnum{
		CREATING: ImageServerStatus{
			value: "CREATING",
		},
		ACTIVE: ImageServerStatus{
			value: "ACTIVE",
		},
		BUILDING: ImageServerStatus{
			value: "BUILDING",
		},
		BUILT: ImageServerStatus{
			value: "BUILT",
		},
		ERROR: ImageServerStatus{
			value: "ERROR",
		},
		DELETING: ImageServerStatus{
			value: "DELETING",
		},
		NULL: ImageServerStatus{
			value: "null",
		},
	}
}

func (c ImageServerStatus) Value() string {
	return c.value
}

func (c ImageServerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageServerStatus) UnmarshalJSON(b []byte) error {
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
