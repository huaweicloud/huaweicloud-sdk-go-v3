package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDirectoriesRequest Request Object
type ListDirectoriesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 获取该目录下的数据，如果有子目录，获取所有子目录的数据
	Type ListDirectoriesRequestType `json:"type"`
}

func (o ListDirectoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectoriesRequest struct{}"
	}

	return strings.Join([]string{"ListDirectoriesRequest", string(data)}, " ")
}

type ListDirectoriesRequestType struct {
	value string
}

type ListDirectoriesRequestTypeEnum struct {
	CODE             ListDirectoriesRequestType
	STANDARD_ELEMENT ListDirectoriesRequestType
}

func GetListDirectoriesRequestTypeEnum() ListDirectoriesRequestTypeEnum {
	return ListDirectoriesRequestTypeEnum{
		CODE: ListDirectoriesRequestType{
			value: "CODE",
		},
		STANDARD_ELEMENT: ListDirectoriesRequestType{
			value: "STANDARD_ELEMENT",
		},
	}
}

func (c ListDirectoriesRequestType) Value() string {
	return c.value
}

func (c ListDirectoriesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDirectoriesRequestType) UnmarshalJSON(b []byte) error {
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
