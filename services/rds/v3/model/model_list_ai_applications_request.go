package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAiApplicationsRequest Request Object
type ListAiApplicationsRequest struct {
	XLanguage *ListAiApplicationsRequestXLanguage `json:"X-Language,omitempty"`

	// 按应用名称查询。
	Name *string `json:"name,omitempty"`

	// 按应用ID查询。
	Id *string `json:"id,omitempty"`

	// 按应用类型查询。Supabase传supabase。
	Type *string `json:"type,omitempty"`

	// 按部署模式查询。
	Mode *string `json:"mode,omitempty"`

	// 按状态查询。
	Status *string `json:"status,omitempty"`

	// 按关联PG实例ID查询。
	ManagedInstanceId *string `json:"managed_instance_id,omitempty"`

	// 按关联PG实例名称查询。
	ManagedInstanceName *string `json:"managed_instance_name,omitempty"`

	// 按是否开启SSL查询。
	EnableSsl *bool `json:"enable_ssl,omitempty"`

	StartCreatedTime *string `json:"start_created_time,omitempty"`

	EndCreatedTime *string `json:"end_created_time,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	SortDir *ListAiApplicationsRequestSortDir `json:"sort_dir,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAiApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListAiApplicationsRequest", string(data)}, " ")
}

type ListAiApplicationsRequestXLanguage struct {
	value string
}

type ListAiApplicationsRequestXLanguageEnum struct {
	ZH_CN ListAiApplicationsRequestXLanguage
	EN_US ListAiApplicationsRequestXLanguage
}

func GetListAiApplicationsRequestXLanguageEnum() ListAiApplicationsRequestXLanguageEnum {
	return ListAiApplicationsRequestXLanguageEnum{
		ZH_CN: ListAiApplicationsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListAiApplicationsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListAiApplicationsRequestXLanguage) Value() string {
	return c.value
}

func (c ListAiApplicationsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAiApplicationsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListAiApplicationsRequestSortDir struct {
	value string
}

type ListAiApplicationsRequestSortDirEnum struct {
	DESC ListAiApplicationsRequestSortDir
	ASC  ListAiApplicationsRequestSortDir
}

func GetListAiApplicationsRequestSortDirEnum() ListAiApplicationsRequestSortDirEnum {
	return ListAiApplicationsRequestSortDirEnum{
		DESC: ListAiApplicationsRequestSortDir{
			value: "desc",
		},
		ASC: ListAiApplicationsRequestSortDir{
			value: "asc",
		},
	}
}

func (c ListAiApplicationsRequestSortDir) Value() string {
	return c.value
}

func (c ListAiApplicationsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAiApplicationsRequestSortDir) UnmarshalJSON(b []byte) error {
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
