package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDataServiceInstancesOverviewRequest Request Object
type ListDataServiceInstancesOverviewRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListDataServiceInstancesOverviewRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType *string `json:"Content-Type,omitempty"`

	// 查询条数限制。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标。
	Offset *int32 `json:"offset,omitempty"`

	// 集群名称。
	Name *string `json:"name,omitempty"`

	// 创建人名称。
	CreateUser *string `json:"create_user,omitempty"`
}

func (o ListDataServiceInstancesOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataServiceInstancesOverviewRequest struct{}"
	}

	return strings.Join([]string{"ListDataServiceInstancesOverviewRequest", string(data)}, " ")
}

type ListDataServiceInstancesOverviewRequestDlmType struct {
	value string
}

type ListDataServiceInstancesOverviewRequestDlmTypeEnum struct {
	SHARED    ListDataServiceInstancesOverviewRequestDlmType
	EXCLUSIVE ListDataServiceInstancesOverviewRequestDlmType
}

func GetListDataServiceInstancesOverviewRequestDlmTypeEnum() ListDataServiceInstancesOverviewRequestDlmTypeEnum {
	return ListDataServiceInstancesOverviewRequestDlmTypeEnum{
		SHARED: ListDataServiceInstancesOverviewRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListDataServiceInstancesOverviewRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListDataServiceInstancesOverviewRequestDlmType) Value() string {
	return c.value
}

func (c ListDataServiceInstancesOverviewRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDataServiceInstancesOverviewRequestDlmType) UnmarshalJSON(b []byte) error {
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
