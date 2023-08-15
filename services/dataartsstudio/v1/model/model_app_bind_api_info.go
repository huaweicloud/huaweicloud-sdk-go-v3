package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppBindApiInfo 应用已绑定的api信息
type AppBindApiInfo struct {

	// API ID
	Id *string `json:"id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API描述
	Description *string `json:"description,omitempty"`

	// 审核时间
	ApprovalTime *int64 `json:"approval_time,omitempty"`

	// API 审核人名称
	Manager *string `json:"manager,omitempty"`

	// 使用截止时间
	Deadline *int64 `json:"deadline,omitempty"`

	// 绑定关系
	RelationshipType *AppBindApiInfoRelationshipType `json:"relationship_type,omitempty"`

	// 静态参数列表
	StaticParams *[]StaticParam `json:"static_params,omitempty"`
}

func (o AppBindApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppBindApiInfo struct{}"
	}

	return strings.Join([]string{"AppBindApiInfo", string(data)}, " ")
}

type AppBindApiInfoRelationshipType struct {
	value string
}

type AppBindApiInfoRelationshipTypeEnum struct {
	LINK_WAITING_CHECK    AppBindApiInfoRelationshipType
	LINKED                AppBindApiInfoRelationshipType
	OFFLINE_WAITING_CHECK AppBindApiInfoRelationshipType
	RENEW_WAITING_CHECK   AppBindApiInfoRelationshipType
}

func GetAppBindApiInfoRelationshipTypeEnum() AppBindApiInfoRelationshipTypeEnum {
	return AppBindApiInfoRelationshipTypeEnum{
		LINK_WAITING_CHECK: AppBindApiInfoRelationshipType{
			value: "LINK_WAITING_CHECK",
		},
		LINKED: AppBindApiInfoRelationshipType{
			value: "LINKED",
		},
		OFFLINE_WAITING_CHECK: AppBindApiInfoRelationshipType{
			value: "OFFLINE_WAITING_CHECK",
		},
		RENEW_WAITING_CHECK: AppBindApiInfoRelationshipType{
			value: "RENEW_WAITING_CHECK",
		},
	}
}

func (c AppBindApiInfoRelationshipType) Value() string {
	return c.value
}

func (c AppBindApiInfoRelationshipType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppBindApiInfoRelationshipType) UnmarshalJSON(b []byte) error {
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
