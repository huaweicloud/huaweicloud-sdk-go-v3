package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScriptsRequest Request Object
type ListScriptsRequest struct {

	// 分页参数：每页返回记录个数限制
	Limit int32 `json:"limit"`

	// 分页参数：上一页最后一个记录id
	Marker *int64 `json:"marker,omitempty"`

	// 脚本名（模糊）
	NameLike *string `json:"name_like,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 风险等级 LOW:低风险 MEDIUM：中风险 HIGH：高风险
	RiskLevel *ListScriptsRequestRiskLevel `json:"risk_level,omitempty"`

	// 脚本类型 SHELL:shell脚本 PYTHON：python脚本 BAT：Bat脚本
	Type *ListScriptsRequestType `json:"type,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 项目ID，一个项目对应一个region
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息
	XUserProfile *string `json:"x-user-profile,omitempty"`
}

func (o ListScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptsRequest", string(data)}, " ")
}

type ListScriptsRequestRiskLevel struct {
	value string
}

type ListScriptsRequestRiskLevelEnum struct {
	LOW    ListScriptsRequestRiskLevel
	MEDIUM ListScriptsRequestRiskLevel
	HIGH   ListScriptsRequestRiskLevel
}

func GetListScriptsRequestRiskLevelEnum() ListScriptsRequestRiskLevelEnum {
	return ListScriptsRequestRiskLevelEnum{
		LOW: ListScriptsRequestRiskLevel{
			value: "LOW",
		},
		MEDIUM: ListScriptsRequestRiskLevel{
			value: "MEDIUM",
		},
		HIGH: ListScriptsRequestRiskLevel{
			value: "HIGH",
		},
	}
}

func (c ListScriptsRequestRiskLevel) Value() string {
	return c.value
}

func (c ListScriptsRequestRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScriptsRequestRiskLevel) UnmarshalJSON(b []byte) error {
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

type ListScriptsRequestType struct {
	value string
}

type ListScriptsRequestTypeEnum struct {
	SHELL  ListScriptsRequestType
	PYTHON ListScriptsRequestType
	BAT    ListScriptsRequestType
}

func GetListScriptsRequestTypeEnum() ListScriptsRequestTypeEnum {
	return ListScriptsRequestTypeEnum{
		SHELL: ListScriptsRequestType{
			value: "SHELL",
		},
		PYTHON: ListScriptsRequestType{
			value: "PYTHON",
		},
		BAT: ListScriptsRequestType{
			value: "BAT",
		},
	}
}

func (c ListScriptsRequestType) Value() string {
	return c.value
}

func (c ListScriptsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScriptsRequestType) UnmarshalJSON(b []byte) error {
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
