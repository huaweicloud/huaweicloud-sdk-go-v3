package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEngineResponse Response Object
type ShowEngineResponse struct {

	// 微服务引擎专享版ID
	Id *string `json:"id,omitempty"`

	// 微服务引擎专享版名称
	Name *string `json:"name,omitempty"`

	// 微服务引擎专享版描述
	Description *string `json:"description,omitempty"`

	// 微服务引擎专享版认证类型
	AuthType *ShowEngineResponseAuthType `json:"authType,omitempty"`

	// 微服务引擎专享版规格
	Flavor *string `json:"flavor,omitempty"`

	// 微服务引擎专享版计费方式
	Payment *string `json:"payment,omitempty"`

	// 微服务引擎专享版当前版本
	Version *string `json:"version,omitempty"`

	// 微服务引擎专享版最新版本
	LatestVersion *string `json:"latestVersion,omitempty"`

	// 微服务引擎专享版状态
	Status *ShowEngineResponseStatus `json:"status,omitempty"`

	// engine 是否是默认引擎
	BeDefault *bool `json:"beDefault,omitempty"`

	// 微服务引擎专享版创建者
	CreateUser *string `json:"createUser,omitempty"`

	// 微服务引擎专享版创建时间
	CreateTime *int64 `json:"createTime,omitempty"`

	CceSpec *Spec `json:"cceSpec,omitempty"`

	ExternalEntrypoint *EngineExternalEntrypoint `json:"externalEntrypoint,omitempty"`

	Reference *EngineReference `json:"reference,omitempty"`

	// 微服务引擎专享版最近的任务ID
	LatestJobId *int32 `json:"latestJobId,omitempty"`

	// 微服务引擎专享版所属企业项目ID
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`

	// 微服务引擎专享版所属企业项目名称
	EnterpriseProjectName *string `json:"enterpriseProjectName,omitempty"`

	// 微服务引擎专享版允许的附加操作
	EngineAdditionalActions *[]ShowEngineResponseEngineAdditionalActions `json:"engineAdditionalActions,omitempty"`

	// 微服务引擎专享版应用部署类型
	SpecType *ShowEngineResponseSpecType `json:"specType,omitempty"`

	// 微服务引擎类型，CSE表示专享版，CSE_Share表示专业版
	Type *ShowEngineResponseType `json:"type,omitempty"`

	// 微服务引擎专享版所属项目ID
	ProjectId *string `json:"projectId,omitempty"`

	// 当前引擎在资源租户侧使用的虚拟机 id 列表
	VmIds          *[]string `json:"vmIds,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineResponse struct{}"
	}

	return strings.Join([]string{"ShowEngineResponse", string(data)}, " ")
}

type ShowEngineResponseAuthType struct {
	value string
}

type ShowEngineResponseAuthTypeEnum struct {
	RBAC ShowEngineResponseAuthType
	NONE ShowEngineResponseAuthType
}

func GetShowEngineResponseAuthTypeEnum() ShowEngineResponseAuthTypeEnum {
	return ShowEngineResponseAuthTypeEnum{
		RBAC: ShowEngineResponseAuthType{
			value: "RBAC",
		},
		NONE: ShowEngineResponseAuthType{
			value: "NONE",
		},
	}
}

func (c ShowEngineResponseAuthType) Value() string {
	return c.value
}

func (c ShowEngineResponseAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseAuthType) UnmarshalJSON(b []byte) error {
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

type ShowEngineResponseStatus struct {
	value string
}

type ShowEngineResponseStatusEnum struct {
	CREATING       ShowEngineResponseStatus
	AVAILABLE      ShowEngineResponseStatus
	UNAVAILABLE    ShowEngineResponseStatus
	DELETING       ShowEngineResponseStatus
	DELETED        ShowEngineResponseStatus
	UPGRADING      ShowEngineResponseStatus
	MODIFYING      ShowEngineResponseStatus
	CREATE_FAILED  ShowEngineResponseStatus
	DELETE_FAILED  ShowEngineResponseStatus
	UPGRADE_FAILED ShowEngineResponseStatus
	MODIFY_FAILED  ShowEngineResponseStatus
	FREEZED        ShowEngineResponseStatus
}

func GetShowEngineResponseStatusEnum() ShowEngineResponseStatusEnum {
	return ShowEngineResponseStatusEnum{
		CREATING: ShowEngineResponseStatus{
			value: "Creating",
		},
		AVAILABLE: ShowEngineResponseStatus{
			value: "Available",
		},
		UNAVAILABLE: ShowEngineResponseStatus{
			value: "Unavailable",
		},
		DELETING: ShowEngineResponseStatus{
			value: "Deleting",
		},
		DELETED: ShowEngineResponseStatus{
			value: "Deleted",
		},
		UPGRADING: ShowEngineResponseStatus{
			value: "Upgrading",
		},
		MODIFYING: ShowEngineResponseStatus{
			value: "Modifying",
		},
		CREATE_FAILED: ShowEngineResponseStatus{
			value: "CreateFailed",
		},
		DELETE_FAILED: ShowEngineResponseStatus{
			value: "DeleteFailed",
		},
		UPGRADE_FAILED: ShowEngineResponseStatus{
			value: "UpgradeFailed",
		},
		MODIFY_FAILED: ShowEngineResponseStatus{
			value: "ModifyFailed",
		},
		FREEZED: ShowEngineResponseStatus{
			value: "Freezed",
		},
	}
}

func (c ShowEngineResponseStatus) Value() string {
	return c.value
}

func (c ShowEngineResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowEngineResponseEngineAdditionalActions struct {
	value string
}

type ShowEngineResponseEngineAdditionalActionsEnum struct {
	NOTING       ShowEngineResponseEngineAdditionalActions
	FORCE_DELETE ShowEngineResponseEngineAdditionalActions
	ROLLBACK     ShowEngineResponseEngineAdditionalActions
	RETRY        ShowEngineResponseEngineAdditionalActions
}

func GetShowEngineResponseEngineAdditionalActionsEnum() ShowEngineResponseEngineAdditionalActionsEnum {
	return ShowEngineResponseEngineAdditionalActionsEnum{
		NOTING: ShowEngineResponseEngineAdditionalActions{
			value: "Noting",
		},
		FORCE_DELETE: ShowEngineResponseEngineAdditionalActions{
			value: "ForceDelete",
		},
		ROLLBACK: ShowEngineResponseEngineAdditionalActions{
			value: "Rollback",
		},
		RETRY: ShowEngineResponseEngineAdditionalActions{
			value: "Retry",
		},
	}
}

func (c ShowEngineResponseEngineAdditionalActions) Value() string {
	return c.value
}

func (c ShowEngineResponseEngineAdditionalActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseEngineAdditionalActions) UnmarshalJSON(b []byte) error {
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

type ShowEngineResponseSpecType struct {
	value string
}

type ShowEngineResponseSpecTypeEnum struct {
	CCE          ShowEngineResponseSpecType
	CSE          ShowEngineResponseSpecType
	SPRING_CLOUD ShowEngineResponseSpecType
}

func GetShowEngineResponseSpecTypeEnum() ShowEngineResponseSpecTypeEnum {
	return ShowEngineResponseSpecTypeEnum{
		CCE: ShowEngineResponseSpecType{
			value: "CCE",
		},
		CSE: ShowEngineResponseSpecType{
			value: "CSE",
		},
		SPRING_CLOUD: ShowEngineResponseSpecType{
			value: "SpringCloud",
		},
	}
}

func (c ShowEngineResponseSpecType) Value() string {
	return c.value
}

func (c ShowEngineResponseSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseSpecType) UnmarshalJSON(b []byte) error {
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

type ShowEngineResponseType struct {
	value string
}

type ShowEngineResponseTypeEnum struct {
	CSE       ShowEngineResponseType
	CSE_SHARE ShowEngineResponseType
}

func GetShowEngineResponseTypeEnum() ShowEngineResponseTypeEnum {
	return ShowEngineResponseTypeEnum{
		CSE: ShowEngineResponseType{
			value: "CSE",
		},
		CSE_SHARE: ShowEngineResponseType{
			value: "CSE_Share",
		},
	}
}

func (c ShowEngineResponseType) Value() string {
	return c.value
}

func (c ShowEngineResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseType) UnmarshalJSON(b []byte) error {
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
