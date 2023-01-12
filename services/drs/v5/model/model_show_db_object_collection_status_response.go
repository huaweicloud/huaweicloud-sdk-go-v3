package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowDbObjectCollectionStatusResponse struct {
	TargetRootDb *TargetRootDb `json:"target_root_db,omitempty"`

	// 数据库对象迁移或同步信息。
	ObjectInfo map[string]DatabaseObject `json:"object_info,omitempty"`

	// 库下表数量的阈值。
	MaxTableNum *int32 `json:"max_table_num,omitempty"`

	// 获取提交查询对象选择信息的状态
	Status *ShowDbObjectCollectionStatusResponseStatus `json:"status,omitempty"`

	// 任务id
	Id *string `json:"id,omitempty"`

	// 该数据库在实时同步场景下的类型
	ObjectScope    *string `json:"object_scope,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDbObjectCollectionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectCollectionStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowDbObjectCollectionStatusResponse", string(data)}, " ")
}

type ShowDbObjectCollectionStatusResponseStatus struct {
	value string
}

type ShowDbObjectCollectionStatusResponseStatusEnum struct {
	SUCCESS ShowDbObjectCollectionStatusResponseStatus
	FAILED  ShowDbObjectCollectionStatusResponseStatus
	PENDING ShowDbObjectCollectionStatusResponseStatus
}

func GetShowDbObjectCollectionStatusResponseStatusEnum() ShowDbObjectCollectionStatusResponseStatusEnum {
	return ShowDbObjectCollectionStatusResponseStatusEnum{
		SUCCESS: ShowDbObjectCollectionStatusResponseStatus{
			value: "success",
		},
		FAILED: ShowDbObjectCollectionStatusResponseStatus{
			value: " failed",
		},
		PENDING: ShowDbObjectCollectionStatusResponseStatus{
			value: " pending",
		},
	}
}

func (c ShowDbObjectCollectionStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowDbObjectCollectionStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbObjectCollectionStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
