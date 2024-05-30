package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEngineResponse Response Object
type ShowEngineResponse struct {

	// 微服务引擎ID
	Id *string `json:"id,omitempty"`

	// 微服务引擎名称
	Name *string `json:"name,omitempty"`

	// 微服务引擎描述
	Description *string `json:"description,omitempty"`

	// 微服务引擎认证类型
	AuthType *ShowEngineResponseAuthType `json:"authType,omitempty"`

	// 微服务引擎规格
	Flavor *string `json:"flavor,omitempty"`

	// 微服务引擎计费方式
	Payment *string `json:"payment,omitempty"`

	// 微服务引擎当前版本
	Version *string `json:"version,omitempty"`

	// 微服务引擎最新版本
	LatestVersion *string `json:"latestVersion,omitempty"`

	// 微服务引擎状态
	Status *ShowEngineResponseStatus `json:"status,omitempty"`

	// engine 是否是默认引擎
	BeDefault *bool `json:"beDefault,omitempty"`

	// 微服务引擎创建者
	CreateUser *string `json:"createUser,omitempty"`

	// 微服务引擎创建时间
	CreateTime *int64 `json:"createTime,omitempty"`

	CceSpec *Spec `json:"cceSpec,omitempty"`

	ExternalEntrypoint *EngineExternalEntrypoint `json:"externalEntrypoint,omitempty"`

	Reference *EngineReference `json:"reference,omitempty"`

	// 微服务引擎最近的任务ID
	LatestJobId *int32 `json:"latestJobId,omitempty"`

	// 微服务引擎所属企业项目ID
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`

	// 微服务引擎所属企业项目名称
	EnterpriseProjectName *string `json:"enterpriseProjectName,omitempty"`

	// 微服务引擎允许的附加操作
	EngineAdditionalActions *[]ShowEngineResponseEngineAdditionalActions `json:"engineAdditionalActions,omitempty"`

	// 微服务引擎应用部署类型
	SpecType *ShowEngineResponseSpecType `json:"specType,omitempty"`

	// 微服务引擎类型，CSE表示专享版，NACOS表示注册配置中心，MICROGATEWAY表示网关
	Type *ShowEngineResponseType `json:"type,omitempty"`

	// 微服务引擎所属项目ID
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
	CSE2          ShowEngineResponseSpecType
	NACOS2        ShowEngineResponseSpecType
	MICRO_GATEWAY ShowEngineResponseSpecType
}

func GetShowEngineResponseSpecTypeEnum() ShowEngineResponseSpecTypeEnum {
	return ShowEngineResponseSpecTypeEnum{
		CSE2: ShowEngineResponseSpecType{
			value: "CSE2",
		},
		NACOS2: ShowEngineResponseSpecType{
			value: "Nacos2",
		},
		MICRO_GATEWAY: ShowEngineResponseSpecType{
			value: "MicroGateway",
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
	CSE          ShowEngineResponseType
	NACOS        ShowEngineResponseType
	MICROGATEWAY ShowEngineResponseType
}

func GetShowEngineResponseTypeEnum() ShowEngineResponseTypeEnum {
	return ShowEngineResponseTypeEnum{
		CSE: ShowEngineResponseType{
			value: "CSE",
		},
		NACOS: ShowEngineResponseType{
			value: "NACOS",
		},
		MICROGATEWAY: ShowEngineResponseType{
			value: "MICROGATEWAY",
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
