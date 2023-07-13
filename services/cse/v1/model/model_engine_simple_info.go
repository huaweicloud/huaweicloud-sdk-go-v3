package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EngineSimpleInfo 查询微服务引擎列表引擎信息
type EngineSimpleInfo struct {

	// 微服务引擎专享版的ID
	Id *string `json:"id,omitempty"`

	// 引擎的名称
	Name *string `json:"name,omitempty"`

	// 微服务引擎专享版所属企业项目ID
	EnterprisProjectId *string `json:"enterpris_project_id,omitempty"`

	// 微服务引擎专享版所属企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 微服务引擎专享版的类型，CSE为专享版引擎，CSE_Share表示为专业版引擎
	Type *EngineSimpleInfoType `json:"type,omitempty"`

	// 微服务引擎专享版的描述
	Description *string `json:"description,omitempty"`

	// 微服务引擎专享版的规格
	Flavor *EngineSimpleInfoFlavor `json:"flavor,omitempty"`

	// 微服务引擎专享版的计费方式，0表示包周期，1表示按需，2表示免费
	Payment *string `json:"payment,omitempty"`

	// 微服务引擎专享版的认证方式，RBAC/NONE
	AuthType *EngineSimpleInfoAuthType `json:"auth_type,omitempty"`

	// 微服务引擎专享版当前的状态
	Status *EngineSimpleInfoStatus `json:"status,omitempty"`

	// 微服务引擎专享版暴露的IP地址
	ExternalAddress *string `json:"external_address,omitempty"`

	// 微服务引擎专享版组件的访问地址。
	ServiceEndpoint map[string]EntrypointItem `json:"service_endpoint,omitempty"`

	// 微服务引擎专享版的公网IP地址
	PublicAddress *string `json:"public_address,omitempty"`

	// 微服务引擎专享版的公网接入地址
	PublicServiceEndpoint map[string]EntrypointItem `json:"public_service_endpoint,omitempty"`

	// 微服务引擎专享版可支持的实例总数
	TotalInstance *int32 `json:"total_instance,omitempty"`

	// 已使用的实例总数
	UsedInstance *int32 `json:"used_instance,omitempty"`

	// 可用实例总数
	AvailableInstance *int32 `json:"available_instance,omitempty"`

	// 微服务引擎专享版当前版本
	Version *string `json:"version,omitempty"`

	// 微服务引擎专享版最新版本
	LatestVersion *string `json:"latest_version,omitempty"`

	// 微服务引擎专享版创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 微服务引擎专享版到期时间
	DueTo *int64 `json:"due_to,omitempty"`

	// 微服务引擎专享版最近的任务ID
	LatestJobId *int32 `json:"latest_job_id,omitempty"`

	// 微服务引擎专享版允许的附加操作
	EngineAdditionalActions *[]EngineSimpleInfoEngineAdditionalActions `json:"engine_additional_actions,omitempty"`

	// 微服务引擎专享版应用部署类型
	SpecType *EngineSimpleInfoSpecType `json:"spec_type,omitempty"`

	Reference *EngineReference `json:"reference,omitempty"`
}

func (o EngineSimpleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineSimpleInfo struct{}"
	}

	return strings.Join([]string{"EngineSimpleInfo", string(data)}, " ")
}

type EngineSimpleInfoType struct {
	value string
}

type EngineSimpleInfoTypeEnum struct {
	CSE       EngineSimpleInfoType
	CSE_SHARE EngineSimpleInfoType
}

func GetEngineSimpleInfoTypeEnum() EngineSimpleInfoTypeEnum {
	return EngineSimpleInfoTypeEnum{
		CSE: EngineSimpleInfoType{
			value: "CSE",
		},
		CSE_SHARE: EngineSimpleInfoType{
			value: "CSE_Share",
		},
	}
}

func (c EngineSimpleInfoType) Value() string {
	return c.value
}

func (c EngineSimpleInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineSimpleInfoType) UnmarshalJSON(b []byte) error {
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

type EngineSimpleInfoFlavor struct {
	value string
}

type EngineSimpleInfoFlavorEnum struct {
	CSE_S1_SMALL2  EngineSimpleInfoFlavor
	CSE_S1_MEDIUM2 EngineSimpleInfoFlavor
	CSE_S1_LARGE2  EngineSimpleInfoFlavor
	CSE_S1_XLARGE2 EngineSimpleInfoFlavor
}

func GetEngineSimpleInfoFlavorEnum() EngineSimpleInfoFlavorEnum {
	return EngineSimpleInfoFlavorEnum{
		CSE_S1_SMALL2: EngineSimpleInfoFlavor{
			value: "cse.s1.small2",
		},
		CSE_S1_MEDIUM2: EngineSimpleInfoFlavor{
			value: "cse.s1.medium2",
		},
		CSE_S1_LARGE2: EngineSimpleInfoFlavor{
			value: "cse.s1.large2",
		},
		CSE_S1_XLARGE2: EngineSimpleInfoFlavor{
			value: "cse.s1.xlarge2",
		},
	}
}

func (c EngineSimpleInfoFlavor) Value() string {
	return c.value
}

func (c EngineSimpleInfoFlavor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineSimpleInfoFlavor) UnmarshalJSON(b []byte) error {
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

type EngineSimpleInfoAuthType struct {
	value string
}

type EngineSimpleInfoAuthTypeEnum struct {
	RBAC EngineSimpleInfoAuthType
	NONE EngineSimpleInfoAuthType
}

func GetEngineSimpleInfoAuthTypeEnum() EngineSimpleInfoAuthTypeEnum {
	return EngineSimpleInfoAuthTypeEnum{
		RBAC: EngineSimpleInfoAuthType{
			value: "RBAC",
		},
		NONE: EngineSimpleInfoAuthType{
			value: "NONE",
		},
	}
}

func (c EngineSimpleInfoAuthType) Value() string {
	return c.value
}

func (c EngineSimpleInfoAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineSimpleInfoAuthType) UnmarshalJSON(b []byte) error {
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

type EngineSimpleInfoStatus struct {
	value string
}

type EngineSimpleInfoStatusEnum struct {
	CREATING       EngineSimpleInfoStatus
	AVAILABLE      EngineSimpleInfoStatus
	UNAVAILABLE    EngineSimpleInfoStatus
	DELETING       EngineSimpleInfoStatus
	DELETED        EngineSimpleInfoStatus
	UPGRADING      EngineSimpleInfoStatus
	MODIFYING      EngineSimpleInfoStatus
	CREATE_FAILED  EngineSimpleInfoStatus
	DELETE_FAILED  EngineSimpleInfoStatus
	UPGRADE_FAILED EngineSimpleInfoStatus
	MODIFY_FAILED  EngineSimpleInfoStatus
	FREEZED        EngineSimpleInfoStatus
}

func GetEngineSimpleInfoStatusEnum() EngineSimpleInfoStatusEnum {
	return EngineSimpleInfoStatusEnum{
		CREATING: EngineSimpleInfoStatus{
			value: "Creating",
		},
		AVAILABLE: EngineSimpleInfoStatus{
			value: "Available",
		},
		UNAVAILABLE: EngineSimpleInfoStatus{
			value: "Unavailable",
		},
		DELETING: EngineSimpleInfoStatus{
			value: "Deleting",
		},
		DELETED: EngineSimpleInfoStatus{
			value: "Deleted",
		},
		UPGRADING: EngineSimpleInfoStatus{
			value: "Upgrading",
		},
		MODIFYING: EngineSimpleInfoStatus{
			value: "Modifying",
		},
		CREATE_FAILED: EngineSimpleInfoStatus{
			value: "CreateFailed",
		},
		DELETE_FAILED: EngineSimpleInfoStatus{
			value: "DeleteFailed",
		},
		UPGRADE_FAILED: EngineSimpleInfoStatus{
			value: "UpgradeFailed",
		},
		MODIFY_FAILED: EngineSimpleInfoStatus{
			value: "ModifyFailed",
		},
		FREEZED: EngineSimpleInfoStatus{
			value: "Freezed",
		},
	}
}

func (c EngineSimpleInfoStatus) Value() string {
	return c.value
}

func (c EngineSimpleInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineSimpleInfoStatus) UnmarshalJSON(b []byte) error {
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

type EngineSimpleInfoEngineAdditionalActions struct {
	value string
}

type EngineSimpleInfoEngineAdditionalActionsEnum struct {
	NOTING       EngineSimpleInfoEngineAdditionalActions
	FORCE_DELETE EngineSimpleInfoEngineAdditionalActions
	ROLLBACK     EngineSimpleInfoEngineAdditionalActions
	RETRY        EngineSimpleInfoEngineAdditionalActions
}

func GetEngineSimpleInfoEngineAdditionalActionsEnum() EngineSimpleInfoEngineAdditionalActionsEnum {
	return EngineSimpleInfoEngineAdditionalActionsEnum{
		NOTING: EngineSimpleInfoEngineAdditionalActions{
			value: "Noting",
		},
		FORCE_DELETE: EngineSimpleInfoEngineAdditionalActions{
			value: "ForceDelete",
		},
		ROLLBACK: EngineSimpleInfoEngineAdditionalActions{
			value: "Rollback",
		},
		RETRY: EngineSimpleInfoEngineAdditionalActions{
			value: "Retry",
		},
	}
}

func (c EngineSimpleInfoEngineAdditionalActions) Value() string {
	return c.value
}

func (c EngineSimpleInfoEngineAdditionalActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineSimpleInfoEngineAdditionalActions) UnmarshalJSON(b []byte) error {
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

type EngineSimpleInfoSpecType struct {
	value string
}

type EngineSimpleInfoSpecTypeEnum struct {
	CCE          EngineSimpleInfoSpecType
	CSE          EngineSimpleInfoSpecType
	SPRING_CLOUD EngineSimpleInfoSpecType
}

func GetEngineSimpleInfoSpecTypeEnum() EngineSimpleInfoSpecTypeEnum {
	return EngineSimpleInfoSpecTypeEnum{
		CCE: EngineSimpleInfoSpecType{
			value: "CCE",
		},
		CSE: EngineSimpleInfoSpecType{
			value: "CSE",
		},
		SPRING_CLOUD: EngineSimpleInfoSpecType{
			value: "SpringCloud",
		},
	}
}

func (c EngineSimpleInfoSpecType) Value() string {
	return c.value
}

func (c EngineSimpleInfoSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineSimpleInfoSpecType) UnmarshalJSON(b []byte) error {
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
