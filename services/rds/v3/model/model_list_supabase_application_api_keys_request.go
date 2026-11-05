package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupabaseApplicationApiKeysRequest Request Object
type ListSupabaseApplicationApiKeysRequest struct {
	XLanguage *ListSupabaseApplicationApiKeysRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`

	Name *string `json:"name,omitempty"`

	Id *string `json:"id,omitempty"`

	Status *ListSupabaseApplicationApiKeysRequestStatus `json:"status,omitempty"`

	Type *ListSupabaseApplicationApiKeysRequestType `json:"type,omitempty"`

	StartCreatedTime *string `json:"start_created_time,omitempty"`

	EndCreatedTime *string `json:"end_created_time,omitempty"`

	SortKey *ListSupabaseApplicationApiKeysRequestSortKey `json:"sort_key,omitempty"`

	SortDir *ListSupabaseApplicationApiKeysRequestSortDir `json:"sort_dir,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSupabaseApplicationApiKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupabaseApplicationApiKeysRequest struct{}"
	}

	return strings.Join([]string{"ListSupabaseApplicationApiKeysRequest", string(data)}, " ")
}

type ListSupabaseApplicationApiKeysRequestXLanguage struct {
	value string
}

type ListSupabaseApplicationApiKeysRequestXLanguageEnum struct {
	ZH_CN ListSupabaseApplicationApiKeysRequestXLanguage
	EN_US ListSupabaseApplicationApiKeysRequestXLanguage
}

func GetListSupabaseApplicationApiKeysRequestXLanguageEnum() ListSupabaseApplicationApiKeysRequestXLanguageEnum {
	return ListSupabaseApplicationApiKeysRequestXLanguageEnum{
		ZH_CN: ListSupabaseApplicationApiKeysRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSupabaseApplicationApiKeysRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSupabaseApplicationApiKeysRequestXLanguage) Value() string {
	return c.value
}

func (c ListSupabaseApplicationApiKeysRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupabaseApplicationApiKeysRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListSupabaseApplicationApiKeysRequestStatus struct {
	value string
}

type ListSupabaseApplicationApiKeysRequestStatusEnum struct {
	NORMAL     ListSupabaseApplicationApiKeysRequestStatus
	CREATING   ListSupabaseApplicationApiKeysRequestStatus
	CREATEFAIL ListSupabaseApplicationApiKeysRequestStatus
	DELETED    ListSupabaseApplicationApiKeysRequestStatus
}

func GetListSupabaseApplicationApiKeysRequestStatusEnum() ListSupabaseApplicationApiKeysRequestStatusEnum {
	return ListSupabaseApplicationApiKeysRequestStatusEnum{
		NORMAL: ListSupabaseApplicationApiKeysRequestStatus{
			value: "normal",
		},
		CREATING: ListSupabaseApplicationApiKeysRequestStatus{
			value: "creating",
		},
		CREATEFAIL: ListSupabaseApplicationApiKeysRequestStatus{
			value: "createfail",
		},
		DELETED: ListSupabaseApplicationApiKeysRequestStatus{
			value: "deleted",
		},
	}
}

func (c ListSupabaseApplicationApiKeysRequestStatus) Value() string {
	return c.value
}

func (c ListSupabaseApplicationApiKeysRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupabaseApplicationApiKeysRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListSupabaseApplicationApiKeysRequestType struct {
	value string
}

type ListSupabaseApplicationApiKeysRequestTypeEnum struct {
	PUBLISHABLE  ListSupabaseApplicationApiKeysRequestType
	SECRET       ListSupabaseApplicationApiKeysRequestType
	ANON         ListSupabaseApplicationApiKeysRequestType
	SERVICE_ROLE ListSupabaseApplicationApiKeysRequestType
}

func GetListSupabaseApplicationApiKeysRequestTypeEnum() ListSupabaseApplicationApiKeysRequestTypeEnum {
	return ListSupabaseApplicationApiKeysRequestTypeEnum{
		PUBLISHABLE: ListSupabaseApplicationApiKeysRequestType{
			value: "publishable",
		},
		SECRET: ListSupabaseApplicationApiKeysRequestType{
			value: "secret",
		},
		ANON: ListSupabaseApplicationApiKeysRequestType{
			value: "anon",
		},
		SERVICE_ROLE: ListSupabaseApplicationApiKeysRequestType{
			value: "service_role",
		},
	}
}

func (c ListSupabaseApplicationApiKeysRequestType) Value() string {
	return c.value
}

func (c ListSupabaseApplicationApiKeysRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupabaseApplicationApiKeysRequestType) UnmarshalJSON(b []byte) error {
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

type ListSupabaseApplicationApiKeysRequestSortKey struct {
	value string
}

type ListSupabaseApplicationApiKeysRequestSortKeyEnum struct {
	NAME         ListSupabaseApplicationApiKeysRequestSortKey
	STATUS       ListSupabaseApplicationApiKeysRequestSortKey
	TYPE         ListSupabaseApplicationApiKeysRequestSortKey
	CREATED_USER ListSupabaseApplicationApiKeysRequestSortKey
	UPDATED_USER ListSupabaseApplicationApiKeysRequestSortKey
	CREATED_TIME ListSupabaseApplicationApiKeysRequestSortKey
	UPDATED_TIME ListSupabaseApplicationApiKeysRequestSortKey
}

func GetListSupabaseApplicationApiKeysRequestSortKeyEnum() ListSupabaseApplicationApiKeysRequestSortKeyEnum {
	return ListSupabaseApplicationApiKeysRequestSortKeyEnum{
		NAME: ListSupabaseApplicationApiKeysRequestSortKey{
			value: "name",
		},
		STATUS: ListSupabaseApplicationApiKeysRequestSortKey{
			value: "status",
		},
		TYPE: ListSupabaseApplicationApiKeysRequestSortKey{
			value: "type",
		},
		CREATED_USER: ListSupabaseApplicationApiKeysRequestSortKey{
			value: "created_user",
		},
		UPDATED_USER: ListSupabaseApplicationApiKeysRequestSortKey{
			value: "updated_user",
		},
		CREATED_TIME: ListSupabaseApplicationApiKeysRequestSortKey{
			value: "created_time",
		},
		UPDATED_TIME: ListSupabaseApplicationApiKeysRequestSortKey{
			value: "updated_time",
		},
	}
}

func (c ListSupabaseApplicationApiKeysRequestSortKey) Value() string {
	return c.value
}

func (c ListSupabaseApplicationApiKeysRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupabaseApplicationApiKeysRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListSupabaseApplicationApiKeysRequestSortDir struct {
	value string
}

type ListSupabaseApplicationApiKeysRequestSortDirEnum struct {
	DESC ListSupabaseApplicationApiKeysRequestSortDir
	ASC  ListSupabaseApplicationApiKeysRequestSortDir
}

func GetListSupabaseApplicationApiKeysRequestSortDirEnum() ListSupabaseApplicationApiKeysRequestSortDirEnum {
	return ListSupabaseApplicationApiKeysRequestSortDirEnum{
		DESC: ListSupabaseApplicationApiKeysRequestSortDir{
			value: "desc",
		},
		ASC: ListSupabaseApplicationApiKeysRequestSortDir{
			value: "asc",
		},
	}
}

func (c ListSupabaseApplicationApiKeysRequestSortDir) Value() string {
	return c.value
}

func (c ListSupabaseApplicationApiKeysRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupabaseApplicationApiKeysRequestSortDir) UnmarshalJSON(b []byte) error {
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
