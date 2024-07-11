package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHostGroupBaseInfosRequest Request Object
type ListHostGroupBaseInfosRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 操作系统：windows|linux
	Os *ListHostGroupBaseInfosRequestOs `json:"os,omitempty"`

	// 分页页码
	PageIndex *int32 `json:"page_index,omitempty"`

	// 分页查询每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 按主机集群名称搜索关键字
	Name *string `json:"name,omitempty"`
}

func (o ListHostGroupBaseInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupBaseInfosRequest struct{}"
	}

	return strings.Join([]string{"ListHostGroupBaseInfosRequest", string(data)}, " ")
}

type ListHostGroupBaseInfosRequestOs struct {
	value string
}

type ListHostGroupBaseInfosRequestOsEnum struct {
	LINUX   ListHostGroupBaseInfosRequestOs
	WINDOWS ListHostGroupBaseInfosRequestOs
}

func GetListHostGroupBaseInfosRequestOsEnum() ListHostGroupBaseInfosRequestOsEnum {
	return ListHostGroupBaseInfosRequestOsEnum{
		LINUX: ListHostGroupBaseInfosRequestOs{
			value: "linux",
		},
		WINDOWS: ListHostGroupBaseInfosRequestOs{
			value: "windows",
		},
	}
}

func (c ListHostGroupBaseInfosRequestOs) Value() string {
	return c.value
}

func (c ListHostGroupBaseInfosRequestOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostGroupBaseInfosRequestOs) UnmarshalJSON(b []byte) error {
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
