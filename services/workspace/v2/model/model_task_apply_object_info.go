package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskApplyObjectInfo 任务应用对象信息
type TaskApplyObjectInfo struct {

	// 应用对象类型，包括DESKTOP（单桌面）、ALL_DESKTOPS（全部桌面）、DESKTOP_POOL（桌面池）、DESKTOP_TAG（桌面标签）、ALL_USERS（全部用户）、USER（单个用户）、USER_GROUP（用户组）、ALL_IMAGES（全部镜像）、APPLICATION_SERVER（应用服务器）、APPLICATION_SERVER_GROUP（应用服务器组）
	ObjectType TaskApplyObjectInfoObjectType `json:"object_type"`

	// 对象ID（object_type为ALL_DESKTOPS、ALL_USERS或ALL_IMAGES时可为null）
	ObjectId *string `json:"object_id,omitempty"`
}

func (o TaskApplyObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskApplyObjectInfo struct{}"
	}

	return strings.Join([]string{"TaskApplyObjectInfo", string(data)}, " ")
}

type TaskApplyObjectInfoObjectType struct {
	value string
}

type TaskApplyObjectInfoObjectTypeEnum struct {
	DESKTOP                  TaskApplyObjectInfoObjectType
	ALL_DESKTOPS             TaskApplyObjectInfoObjectType
	DESKTOP_POOL             TaskApplyObjectInfoObjectType
	DESKTOP_TAG              TaskApplyObjectInfoObjectType
	ALL_IMAGES               TaskApplyObjectInfoObjectType
	APPLICATION_SERVER       TaskApplyObjectInfoObjectType
	APPLICATION_SERVER_GROUP TaskApplyObjectInfoObjectType
}

func GetTaskApplyObjectInfoObjectTypeEnum() TaskApplyObjectInfoObjectTypeEnum {
	return TaskApplyObjectInfoObjectTypeEnum{
		DESKTOP: TaskApplyObjectInfoObjectType{
			value: "DESKTOP",
		},
		ALL_DESKTOPS: TaskApplyObjectInfoObjectType{
			value: "ALL_DESKTOPS",
		},
		DESKTOP_POOL: TaskApplyObjectInfoObjectType{
			value: "DESKTOP_POOL",
		},
		DESKTOP_TAG: TaskApplyObjectInfoObjectType{
			value: "DESKTOP_TAG",
		},
		ALL_IMAGES: TaskApplyObjectInfoObjectType{
			value: "ALL_IMAGES",
		},
		APPLICATION_SERVER: TaskApplyObjectInfoObjectType{
			value: "APPLICATION_SERVER",
		},
		APPLICATION_SERVER_GROUP: TaskApplyObjectInfoObjectType{
			value: "APPLICATION_SERVER_GROUP",
		},
	}
}

func (c TaskApplyObjectInfoObjectType) Value() string {
	return c.value
}

func (c TaskApplyObjectInfoObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskApplyObjectInfoObjectType) UnmarshalJSON(b []byte) error {
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
