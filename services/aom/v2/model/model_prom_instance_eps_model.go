package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PromInstanceEpsModel struct {

	// 普罗实例名称 名称不能以下划线或中划线开头结尾，只含有中文、英文、数字、下划线、中划线、长度1-100
	PromName string `json:"prom_name"`

	// 普罗实例ID
	PromId *string `json:"prom_id,omitempty"`

	// 普罗实例类型,DEFAULT,ECS,VPC,CCE,REMOTE_WRITE,KUBERNETES,CLOUD_SERVICE,ACROSS_ACCOUNT
	PromType PromInstanceEpsModelPromType `json:"prom_type"`

	// 普罗实例版本号
	PromVersion *string `json:"prom_version,omitempty"`

	// CCE场景特殊字段
	CceSpec *string `json:"cce_spec,omitempty"`

	PromConfig *PromConfigModel `json:"prom_config,omitempty"`

	// 普罗实例创建时间戳
	PromCreateTimestamp *int64 `json:"prom_create_timestamp,omitempty"`

	// 普罗实例更新时间戳
	PromUpdateTimestamp *int64 `json:"prom_update_timestamp,omitempty"`

	// 普罗实例状态 true/false
	PromStatus *string `json:"prom_status,omitempty"`

	// 普罗实例所属的企业项目
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 普罗实例所属projectId
	ProjectId *string `json:"project_id,omitempty"`

	// 删除标记
	IsDeletedTag *int64 `json:"is_deleted_tag,omitempty"`

	// 删除时间
	DeletedTime *int64 `json:"deleted_time,omitempty"`

	PromSpecConfig *PromConfigModel `json:"prom_spec_config,omitempty"`

	// 普罗实例所属CCE特殊配置
	CceSpecConfig *string `json:"cce_spec_config,omitempty"`

	Application *ApplicationModel `json:"application,omitempty"`
}

func (o PromInstanceEpsModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PromInstanceEpsModel struct{}"
	}

	return strings.Join([]string{"PromInstanceEpsModel", string(data)}, " ")
}

type PromInstanceEpsModelPromType struct {
	value string
}

type PromInstanceEpsModelPromTypeEnum struct {
	DEFAULT        PromInstanceEpsModelPromType
	ECS            PromInstanceEpsModelPromType
	VPC            PromInstanceEpsModelPromType
	CCE            PromInstanceEpsModelPromType
	REMOTE_WRITE   PromInstanceEpsModelPromType
	KUBERNETES     PromInstanceEpsModelPromType
	CLOUD_SERVICE  PromInstanceEpsModelPromType
	ACROSS_ACCOUNT PromInstanceEpsModelPromType
}

func GetPromInstanceEpsModelPromTypeEnum() PromInstanceEpsModelPromTypeEnum {
	return PromInstanceEpsModelPromTypeEnum{
		DEFAULT: PromInstanceEpsModelPromType{
			value: "DEFAULT",
		},
		ECS: PromInstanceEpsModelPromType{
			value: "ECS",
		},
		VPC: PromInstanceEpsModelPromType{
			value: "VPC",
		},
		CCE: PromInstanceEpsModelPromType{
			value: "CCE",
		},
		REMOTE_WRITE: PromInstanceEpsModelPromType{
			value: "REMOTE_WRITE",
		},
		KUBERNETES: PromInstanceEpsModelPromType{
			value: "KUBERNETES",
		},
		CLOUD_SERVICE: PromInstanceEpsModelPromType{
			value: "CLOUD_SERVICE",
		},
		ACROSS_ACCOUNT: PromInstanceEpsModelPromType{
			value: "ACROSS_ACCOUNT",
		},
	}
}

func (c PromInstanceEpsModelPromType) Value() string {
	return c.value
}

func (c PromInstanceEpsModelPromType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PromInstanceEpsModelPromType) UnmarshalJSON(b []byte) error {
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
