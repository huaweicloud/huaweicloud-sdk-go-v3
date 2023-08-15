package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RecordForGetAuthApp struct {

	// 应用编号
	AppId *string `json:"app_id,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 集群实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 集群实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 使用截止时间
	ApiUsingTime *int64 `json:"api_using_time,omitempty"`

	// 授权时间
	ApprovalTime *int64 `json:"approval_time,omitempty"`

	// 绑定关系
	RelationshipType *RecordForGetAuthAppRelationshipType `json:"relationship_type,omitempty"`

	// 静态参数列表
	StaticParams *[]StaticParam `json:"static_params,omitempty"`
}

func (o RecordForGetAuthApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordForGetAuthApp struct{}"
	}

	return strings.Join([]string{"RecordForGetAuthApp", string(data)}, " ")
}

type RecordForGetAuthAppRelationshipType struct {
	value string
}

type RecordForGetAuthAppRelationshipTypeEnum struct {
	LINK_WAITING_CHECK    RecordForGetAuthAppRelationshipType
	LINKED                RecordForGetAuthAppRelationshipType
	OFFLINE_WAITING_CHECK RecordForGetAuthAppRelationshipType
	RENEW_WAITING_CHECK   RecordForGetAuthAppRelationshipType
}

func GetRecordForGetAuthAppRelationshipTypeEnum() RecordForGetAuthAppRelationshipTypeEnum {
	return RecordForGetAuthAppRelationshipTypeEnum{
		LINK_WAITING_CHECK: RecordForGetAuthAppRelationshipType{
			value: "LINK_WAITING_CHECK",
		},
		LINKED: RecordForGetAuthAppRelationshipType{
			value: "LINKED",
		},
		OFFLINE_WAITING_CHECK: RecordForGetAuthAppRelationshipType{
			value: "OFFLINE_WAITING_CHECK",
		},
		RENEW_WAITING_CHECK: RecordForGetAuthAppRelationshipType{
			value: "RENEW_WAITING_CHECK",
		},
	}
}

func (c RecordForGetAuthAppRelationshipType) Value() string {
	return c.value
}

func (c RecordForGetAuthAppRelationshipType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordForGetAuthAppRelationshipType) UnmarshalJSON(b []byte) error {
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
