package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProtectableRequest Request Object
type ListProtectableRequest struct {

	// 每页显示的条目数量，每页最多支持50条
	Limit *int32 `json:"limit,omitempty"`

	// 上一次查询最后一条的ID
	Marker *string `json:"marker,omitempty"`

	// 按名称过滤
	Name *string `json:"name,omitempty"`

	// 偏移值
	Offset *int32 `json:"offset,omitempty"`

	// 对象类型
	ProtectableType ListProtectableRequestProtectableType `json:"protectable_type"`

	// 资源的状态，如available，error 等
	Status *string `json:"status,omitempty"`

	// 根据资源id过滤
	Id *string `json:"id,omitempty"`

	// 根据该id过滤属于该服务器的所有磁盘，支持企业多项目的用户才能传入此参数
	ServerId *string `json:"server_id,omitempty"`
}

func (o ListProtectableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectableRequest struct{}"
	}

	return strings.Join([]string{"ListProtectableRequest", string(data)}, " ")
}

type ListProtectableRequestProtectableType struct {
	value string
}

type ListProtectableRequestProtectableTypeEnum struct {
	SERVER ListProtectableRequestProtectableType
	DISK   ListProtectableRequestProtectableType
}

func GetListProtectableRequestProtectableTypeEnum() ListProtectableRequestProtectableTypeEnum {
	return ListProtectableRequestProtectableTypeEnum{
		SERVER: ListProtectableRequestProtectableType{
			value: "server",
		},
		DISK: ListProtectableRequestProtectableType{
			value: "disk",
		},
	}
}

func (c ListProtectableRequestProtectableType) Value() string {
	return c.value
}

func (c ListProtectableRequestProtectableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProtectableRequestProtectableType) UnmarshalJSON(b []byte) error {
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
