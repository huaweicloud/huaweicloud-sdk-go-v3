package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPublicScriptsRequest Request Object
type ListPublicScriptsRequest struct {

	// 分页参数：每页返回记录个数限制
	Limit int32 `json:"limit"`

	// 分页参数：上一页最后一个记录id
	Marker *int64 `json:"marker,omitempty"`

	// 脚本名（只支持右模糊搜索）
	NameLike *string `json:"name_like,omitempty"`

	// 脚本名（精确搜索）
	Name *string `json:"name,omitempty"`

	// 风险等级 LOW：低风险 MEDIUM：中风险 HIGH：高风险
	RiskLevel *ListPublicScriptsRequestRiskLevel `json:"risk_level,omitempty"`

	// 脚本类型 SHELL：shell脚本 PYTHON：python脚本 BAT：bat脚本
	Type *ListPublicScriptsRequestType `json:"type,omitempty"`

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 项目ID，一个项目对应一个region
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息
	XUserProfile *string `json:"x-user-profile,omitempty"`
}

func (o ListPublicScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicScriptsRequest", string(data)}, " ")
}

type ListPublicScriptsRequestRiskLevel struct {
	value string
}

type ListPublicScriptsRequestRiskLevelEnum struct {
	LOW    ListPublicScriptsRequestRiskLevel
	MEDIUM ListPublicScriptsRequestRiskLevel
	HIGH   ListPublicScriptsRequestRiskLevel
}

func GetListPublicScriptsRequestRiskLevelEnum() ListPublicScriptsRequestRiskLevelEnum {
	return ListPublicScriptsRequestRiskLevelEnum{
		LOW: ListPublicScriptsRequestRiskLevel{
			value: "LOW",
		},
		MEDIUM: ListPublicScriptsRequestRiskLevel{
			value: "MEDIUM",
		},
		HIGH: ListPublicScriptsRequestRiskLevel{
			value: "HIGH",
		},
	}
}

func (c ListPublicScriptsRequestRiskLevel) Value() string {
	return c.value
}

func (c ListPublicScriptsRequestRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicScriptsRequestRiskLevel) UnmarshalJSON(b []byte) error {
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

type ListPublicScriptsRequestType struct {
	value string
}

type ListPublicScriptsRequestTypeEnum struct {
	SHELL  ListPublicScriptsRequestType
	PYTHON ListPublicScriptsRequestType
	BAT    ListPublicScriptsRequestType
}

func GetListPublicScriptsRequestTypeEnum() ListPublicScriptsRequestTypeEnum {
	return ListPublicScriptsRequestTypeEnum{
		SHELL: ListPublicScriptsRequestType{
			value: "SHELL",
		},
		PYTHON: ListPublicScriptsRequestType{
			value: "PYTHON",
		},
		BAT: ListPublicScriptsRequestType{
			value: "BAT",
		},
	}
}

func (c ListPublicScriptsRequestType) Value() string {
	return c.value
}

func (c ListPublicScriptsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicScriptsRequestType) UnmarshalJSON(b []byte) error {
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
