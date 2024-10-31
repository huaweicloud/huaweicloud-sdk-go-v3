package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServicesInstancesRequest Request Object
type ListServicesInstancesRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 通过service Instance id检索，32~36位的英文、数字、短横组合
	Id *string `json:"id,omitempty"`

	// 通过名字搜索Service Instance，支持模糊查询
	Name *string `json:"name,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 通过id检索Endpoint的参数
	EndpointId *string `json:"endpoint_id,omitempty"`

	// Service ID或者Model ID
	SourceId *string `json:"source_id,omitempty"`

	// Service version ID或者Model version ID
	VersionId *string `json:"version_id,omitempty"`

	// Service的类型，可选值： - PGSQL_SERVICE：DWS Pay-By-Query - LLM_MODEL：大语言模型
	Type *ListServicesInstancesRequestType `json:"type,omitempty"`

	// 可见性检索的参数，可选值为： - PRIVATE: 私有，用户自己创建的； - PUBLIC:公共，查询所有公共的，包括其他用户创建的； - 默认为空，不填表示不限制，则查出当前用户下的，包括PRIVATE和PUBLIC，不包括其他用户创建的。
	Visibility *string `json:"visibility,omitempty"`

	// 根据字段排序，可选值： - CREATE_TIME：创建时间 - DURATION: 运行时间
	SortBy *ListServicesInstancesRequestSortBy `json:"sort_by,omitempty"`

	// 排序方式，可选值： - ASC：正序排序 - DESC: 倒序排序
	OrderBy *ListServicesInstancesRequestOrderBy `json:"order_by,omitempty"`
}

func (o ListServicesInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicesInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListServicesInstancesRequest", string(data)}, " ")
}

type ListServicesInstancesRequestType struct {
	value string
}

type ListServicesInstancesRequestTypeEnum struct {
	PGSQL_SERVICE ListServicesInstancesRequestType
	LLM_MODEL     ListServicesInstancesRequestType
}

func GetListServicesInstancesRequestTypeEnum() ListServicesInstancesRequestTypeEnum {
	return ListServicesInstancesRequestTypeEnum{
		PGSQL_SERVICE: ListServicesInstancesRequestType{
			value: "PGSQL_SERVICE",
		},
		LLM_MODEL: ListServicesInstancesRequestType{
			value: "LLM_MODEL",
		},
	}
}

func (c ListServicesInstancesRequestType) Value() string {
	return c.value
}

func (c ListServicesInstancesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicesInstancesRequestType) UnmarshalJSON(b []byte) error {
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

type ListServicesInstancesRequestSortBy struct {
	value string
}

type ListServicesInstancesRequestSortByEnum struct {
	CREATE_TIME ListServicesInstancesRequestSortBy
	DURATION    ListServicesInstancesRequestSortBy
}

func GetListServicesInstancesRequestSortByEnum() ListServicesInstancesRequestSortByEnum {
	return ListServicesInstancesRequestSortByEnum{
		CREATE_TIME: ListServicesInstancesRequestSortBy{
			value: "CREATE_TIME",
		},
		DURATION: ListServicesInstancesRequestSortBy{
			value: "DURATION",
		},
	}
}

func (c ListServicesInstancesRequestSortBy) Value() string {
	return c.value
}

func (c ListServicesInstancesRequestSortBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicesInstancesRequestSortBy) UnmarshalJSON(b []byte) error {
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

type ListServicesInstancesRequestOrderBy struct {
	value string
}

type ListServicesInstancesRequestOrderByEnum struct {
	ASC  ListServicesInstancesRequestOrderBy
	DESC ListServicesInstancesRequestOrderBy
}

func GetListServicesInstancesRequestOrderByEnum() ListServicesInstancesRequestOrderByEnum {
	return ListServicesInstancesRequestOrderByEnum{
		ASC: ListServicesInstancesRequestOrderBy{
			value: "ASC",
		},
		DESC: ListServicesInstancesRequestOrderBy{
			value: "DESC",
		},
	}
}

func (c ListServicesInstancesRequestOrderBy) Value() string {
	return c.value
}

func (c ListServicesInstancesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicesInstancesRequestOrderBy) UnmarshalJSON(b []byte) error {
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
