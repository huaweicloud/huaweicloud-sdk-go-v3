package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDbObjectsListResponse Response Object
type ShowDbObjectsListResponse struct {
	TargetRootDb *TargetRootDb `json:"target_root_db,omitempty"`

	// 数据库对象迁移或同步信息。
	ObjectInfo map[string]DatabaseObject `json:"object_info,omitempty"`

	// 库下表数量的阈值。
	MaxTableNum *int32 `json:"max_table_num,omitempty"`

	// 获取提交查询对象选择信息的状态
	Status *ShowDbObjectsListResponseStatus `json:"status,omitempty"`

	// 任务id
	Id *string `json:"id,omitempty"`

	// 该数据库在实时同步场景下的类型
	ObjectScope    *string `json:"object_scope,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDbObjectsListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectsListResponse struct{}"
	}

	return strings.Join([]string{"ShowDbObjectsListResponse", string(data)}, " ")
}

type ShowDbObjectsListResponseStatus struct {
	value string
}

type ShowDbObjectsListResponseStatusEnum struct {
	SUCCESS ShowDbObjectsListResponseStatus
	FAILED  ShowDbObjectsListResponseStatus
	PENDING ShowDbObjectsListResponseStatus
}

func GetShowDbObjectsListResponseStatusEnum() ShowDbObjectsListResponseStatusEnum {
	return ShowDbObjectsListResponseStatusEnum{
		SUCCESS: ShowDbObjectsListResponseStatus{
			value: "success",
		},
		FAILED: ShowDbObjectsListResponseStatus{
			value: " failed",
		},
		PENDING: ShowDbObjectsListResponseStatus{
			value: " pending",
		},
	}
}

func (c ShowDbObjectsListResponseStatus) Value() string {
	return c.value
}

func (c ShowDbObjectsListResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbObjectsListResponseStatus) UnmarshalJSON(b []byte) error {
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
