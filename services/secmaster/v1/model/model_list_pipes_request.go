package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPipesRequest Request Object
type ListPipesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// pipe name
	PipeName *string `json:"pipe_name,omitempty"`

	// 数据管道类型；可选值：system-defined(系统预定义)、user-defined(用户自定义)
	PipeType *ListPipesRequestPipeType `json:"pipe_type,omitempty"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序顺序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListPipesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipesRequest struct{}"
	}

	return strings.Join([]string{"ListPipesRequest", string(data)}, " ")
}

type ListPipesRequestPipeType struct {
	value string
}

type ListPipesRequestPipeTypeEnum struct {
	SYSTEM_DEFINED ListPipesRequestPipeType
	USER_DEFINED   ListPipesRequestPipeType
}

func GetListPipesRequestPipeTypeEnum() ListPipesRequestPipeTypeEnum {
	return ListPipesRequestPipeTypeEnum{
		SYSTEM_DEFINED: ListPipesRequestPipeType{
			value: "system-defined",
		},
		USER_DEFINED: ListPipesRequestPipeType{
			value: "user-defined",
		},
	}
}

func (c ListPipesRequestPipeType) Value() string {
	return c.value
}

func (c ListPipesRequestPipeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPipesRequestPipeType) UnmarshalJSON(b []byte) error {
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
