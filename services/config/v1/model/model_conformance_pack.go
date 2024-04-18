package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConformancePack 合规规则包详情。
type ConformancePack struct {

	// 合规规则包ID。
	Id *string `json:"id,omitempty"`

	// 合规规则包名称。
	Name *string `json:"name,omitempty"`

	// 资源栈（stack）的唯一ID。
	StackId *string `json:"stack_id,omitempty"`

	// 资源栈（stack）的名称。
	StackName *string `json:"stack_name,omitempty"`

	// 部署ID。
	DeploymentId *string `json:"deployment_id,omitempty"`

	// 合规规则包创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 合规规则包更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 预定义合规规则包模板名称。
	TemplateKey *string `json:"template_key,omitempty"`

	// 合规规则包模板OBS地址
	TemplateUri *string `json:"template_uri,omitempty"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	// 合规规则包部署状态。
	Status *ConformancePackStatus `json:"status,omitempty"`

	// 部署或删除合规规则包错误时的错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 合规规则包参数。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// 创建者
	CreatedBy *string `json:"created_by,omitempty"`
}

func (o ConformancePack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConformancePack struct{}"
	}

	return strings.Join([]string{"ConformancePack", string(data)}, " ")
}

type ConformancePackStatus struct {
	value string
}

type ConformancePackStatusEnum struct {
	CREATE_SUCCESSFUL    ConformancePackStatus
	CREATE_IN_PROGRESS   ConformancePackStatus
	CREATE_FAILED        ConformancePackStatus
	DELETE_IN_PROGRESS   ConformancePackStatus
	DELETE_FAILED        ConformancePackStatus
	ROLLBACK_SUCCESSFUL  ConformancePackStatus
	ROLLBACK_IN_PROGRESS ConformancePackStatus
	ROLLBACK_FAILED      ConformancePackStatus
	UPDATE_SUCCESSFUL    ConformancePackStatus
	UPDATE_IN_PROGRESS   ConformancePackStatus
	UPDATE_FAILED        ConformancePackStatus
}

func GetConformancePackStatusEnum() ConformancePackStatusEnum {
	return ConformancePackStatusEnum{
		CREATE_SUCCESSFUL: ConformancePackStatus{
			value: "CREATE_SUCCESSFUL",
		},
		CREATE_IN_PROGRESS: ConformancePackStatus{
			value: "CREATE_IN_PROGRESS",
		},
		CREATE_FAILED: ConformancePackStatus{
			value: "CREATE_FAILED",
		},
		DELETE_IN_PROGRESS: ConformancePackStatus{
			value: "DELETE_IN_PROGRESS",
		},
		DELETE_FAILED: ConformancePackStatus{
			value: "DELETE_FAILED",
		},
		ROLLBACK_SUCCESSFUL: ConformancePackStatus{
			value: "ROLLBACK_SUCCESSFUL",
		},
		ROLLBACK_IN_PROGRESS: ConformancePackStatus{
			value: "ROLLBACK_IN_PROGRESS",
		},
		ROLLBACK_FAILED: ConformancePackStatus{
			value: "ROLLBACK_FAILED",
		},
		UPDATE_SUCCESSFUL: ConformancePackStatus{
			value: "UPDATE_SUCCESSFUL",
		},
		UPDATE_IN_PROGRESS: ConformancePackStatus{
			value: "UPDATE_IN_PROGRESS",
		},
		UPDATE_FAILED: ConformancePackStatus{
			value: "UPDATE_FAILED",
		},
	}
}

func (c ConformancePackStatus) Value() string {
	return c.value
}

func (c ConformancePackStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConformancePackStatus) UnmarshalJSON(b []byte) error {
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
