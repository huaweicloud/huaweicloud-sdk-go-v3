package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListDbObjectsResponse struct {
	TargetRootDb *TargetRootDb `json:"target_root_db,omitempty"`

	// 数据库对象迁移或同步信息。
	ObjectInfo map[string]DatabaseObject `json:"object_info,omitempty"`

	// 库下表数量的阈值。
	MaxTableNum *int32 `json:"max_table_num,omitempty"`

	// 获取提交查询对象选择信息的状态
	Status *ListDbObjectsResponseStatus `json:"status,omitempty"`

	// 任务id
	Id *string `json:"id,omitempty"`

	// 该数据库在实时同步场景下的类型
	ObjectScope    *string `json:"object_scope,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDbObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListDbObjectsResponse", string(data)}, " ")
}

type ListDbObjectsResponseStatus struct {
	value string
}

type ListDbObjectsResponseStatusEnum struct {
	SUCCESS ListDbObjectsResponseStatus
	FAILED  ListDbObjectsResponseStatus
	PENDING ListDbObjectsResponseStatus
}

func GetListDbObjectsResponseStatusEnum() ListDbObjectsResponseStatusEnum {
	return ListDbObjectsResponseStatusEnum{
		SUCCESS: ListDbObjectsResponseStatus{
			value: "success",
		},
		FAILED: ListDbObjectsResponseStatus{
			value: " failed",
		},
		PENDING: ListDbObjectsResponseStatus{
			value: " pending",
		},
	}
}

func (c ListDbObjectsResponseStatus) Value() string {
	return c.value
}

func (c ListDbObjectsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDbObjectsResponseStatus) UnmarshalJSON(b []byte) error {
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
