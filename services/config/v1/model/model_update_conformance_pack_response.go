package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateConformancePackResponse Response Object
type UpdateConformancePackResponse struct {

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
	Status *UpdateConformancePackResponseStatus `json:"status,omitempty"`

	// 部署或删除合规规则包错误时的错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 合规规则包参数。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// 创建者
	CreatedBy      *string `json:"created_by,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateConformancePackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConformancePackResponse struct{}"
	}

	return strings.Join([]string{"UpdateConformancePackResponse", string(data)}, " ")
}

type UpdateConformancePackResponseStatus struct {
	value string
}

type UpdateConformancePackResponseStatusEnum struct {
	CREATE_SUCCESSFUL    UpdateConformancePackResponseStatus
	CREATE_IN_PROGRESS   UpdateConformancePackResponseStatus
	CREATE_FAILED        UpdateConformancePackResponseStatus
	DELETE_IN_PROGRESS   UpdateConformancePackResponseStatus
	DELETE_FAILED        UpdateConformancePackResponseStatus
	ROLLBACK_SUCCESSFUL  UpdateConformancePackResponseStatus
	ROLLBACK_IN_PROGRESS UpdateConformancePackResponseStatus
	ROLLBACK_FAILED      UpdateConformancePackResponseStatus
	UPDATE_SUCCESSFUL    UpdateConformancePackResponseStatus
	UPDATE_IN_PROGRESS   UpdateConformancePackResponseStatus
	UPDATE_FAILED        UpdateConformancePackResponseStatus
}

func GetUpdateConformancePackResponseStatusEnum() UpdateConformancePackResponseStatusEnum {
	return UpdateConformancePackResponseStatusEnum{
		CREATE_SUCCESSFUL: UpdateConformancePackResponseStatus{
			value: "CREATE_SUCCESSFUL",
		},
		CREATE_IN_PROGRESS: UpdateConformancePackResponseStatus{
			value: "CREATE_IN_PROGRESS",
		},
		CREATE_FAILED: UpdateConformancePackResponseStatus{
			value: "CREATE_FAILED",
		},
		DELETE_IN_PROGRESS: UpdateConformancePackResponseStatus{
			value: "DELETE_IN_PROGRESS",
		},
		DELETE_FAILED: UpdateConformancePackResponseStatus{
			value: "DELETE_FAILED",
		},
		ROLLBACK_SUCCESSFUL: UpdateConformancePackResponseStatus{
			value: "ROLLBACK_SUCCESSFUL",
		},
		ROLLBACK_IN_PROGRESS: UpdateConformancePackResponseStatus{
			value: "ROLLBACK_IN_PROGRESS",
		},
		ROLLBACK_FAILED: UpdateConformancePackResponseStatus{
			value: "ROLLBACK_FAILED",
		},
		UPDATE_SUCCESSFUL: UpdateConformancePackResponseStatus{
			value: "UPDATE_SUCCESSFUL",
		},
		UPDATE_IN_PROGRESS: UpdateConformancePackResponseStatus{
			value: "UPDATE_IN_PROGRESS",
		},
		UPDATE_FAILED: UpdateConformancePackResponseStatus{
			value: "UPDATE_FAILED",
		},
	}
}

func (c UpdateConformancePackResponseStatus) Value() string {
	return c.value
}

func (c UpdateConformancePackResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateConformancePackResponseStatus) UnmarshalJSON(b []byte) error {
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
