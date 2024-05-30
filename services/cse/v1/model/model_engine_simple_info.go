package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EngineSimpleInfo 查询微服务引擎列表引擎信息
type EngineSimpleInfo struct {

	// 微服务引擎的ID
	Id *string `json:"id,omitempty"`

	// 引擎的名称
	Name *string `json:"name,omitempty"`

	// 微服务引擎所属企业项目ID
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`

	// 微服务引擎所属企业项目名称
	EnterpriseProjectName *string `json:"enterpriseProjectName,omitempty"`

	// 微服务引擎的类型，CSE为专享版引擎，CSE_Share表示为专业版引擎
	Type *EngineSimpleInfoType `json:"type,omitempty"`

	// 微服务引擎的描述
	Description *string `json:"description,omitempty"`

	// 微服务引擎的规格
	Flavor *EngineSimpleInfoFlavor `json:"flavor,omitempty"`

	// 微服务引擎的计费方式，0表示包周期，1表示按需，2表示免费
	Payment *string `json:"payment,omitempty"`

	// 微服务引擎的认证方式，RBAC/NONE
	AuthType *EngineSimpleInfoAuthType `json:"authType,omitempty"`

	// 微服务引擎当前的状态
	Status *EngineSimpleInfoStatus `json:"status,omitempty"`

	// 微服务引擎暴露的IP地址
	ExternalAddress *string `json:"externalAddress,omitempty"`

	// 微服务引擎组件的访问地址。
	ServiceEndpoint map[string]EntrypointItem `json:"serviceEndpoint,omitempty"`

	// 微服务引擎的公网IP地址
	PublicAddress *string `json:"publicAddress,omitempty"`

	// 微服务引擎的公网接入地址
	PublicServiceEndpoint map[string]EntrypointItem `json:"publicServiceEndpoint,omitempty"`

	// 微服务引擎可支持的实例总数
	TotalInstance *int32 `json:"totalInstance,omitempty"`

	// 已使用的实例总数
	UsedInstance *int32 `json:"usedInstance,omitempty"`

	// 可用实例总数
	AvailableInstance *int32 `json:"availableInstance,omitempty"`

	// 微服务引擎当前版本
	Version *string `json:"version,omitempty"`

	// 微服务引擎最新版本
	LatestVersion *string `json:"latestVersion,omitempty"`

	// 微服务引擎创建时间
	CreateTime *int64 `json:"createTime,omitempty"`

	// 微服务引擎到期时间
	DueTo *int64 `json:"dueTo,omitempty"`

	// 微服务引擎最近的任务ID
	LatestJobId *int32 `json:"latestJobId,omitempty"`

	// 微服务引擎允许的附加操作
	EngineAdditionalActions *[]EngineSimpleInfoEngineAdditionalActions `json:"engineAdditionalActions,omitempty"`

	// 微服务引擎应用部署类型
	SpecType *EngineSimpleInfoSpecType `json:"specType,omitempty"`

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
	CSE_S1_SMALL2                 EngineSimpleInfoFlavor
	CSE_S1_MEDIUM2                EngineSimpleInfoFlavor
	CSE_S1_LARGE2                 EngineSimpleInfoFlavor
	CSE_S1_XLARGE2                EngineSimpleInfoFlavor
	CSE_NACOS2_C1_LARGE_10        EngineSimpleInfoFlavor
	CSE_NACOS2_C1_XLARGE_20       EngineSimpleInfoFlavor
	CSE_NACOS2_C1_XLARGE_50       EngineSimpleInfoFlavor
	CSE_NACOS2_C1_XLARGE_60       EngineSimpleInfoFlavor
	CSE_NACOS2_C1_2XLARGE_100     EngineSimpleInfoFlavor
	CSE_MICROGATEWAY_PRO_SMALL_1  EngineSimpleInfoFlavor
	CSE_MICROGATEWAY_PRO_MEDIUM_1 EngineSimpleInfoFlavor
	CSE_MICROGATEWAY_PRO_LARGE_1  EngineSimpleInfoFlavor
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
		CSE_NACOS2_C1_LARGE_10: EngineSimpleInfoFlavor{
			value: "cse.nacos2.c1.large.10",
		},
		CSE_NACOS2_C1_XLARGE_20: EngineSimpleInfoFlavor{
			value: "cse.nacos2.c1.xlarge.20",
		},
		CSE_NACOS2_C1_XLARGE_50: EngineSimpleInfoFlavor{
			value: "cse.nacos2.c1.xlarge.50",
		},
		CSE_NACOS2_C1_XLARGE_60: EngineSimpleInfoFlavor{
			value: "cse.nacos2.c1.xlarge.60",
		},
		CSE_NACOS2_C1_2XLARGE_100: EngineSimpleInfoFlavor{
			value: "cse.nacos2.c1.2xlarge.100",
		},
		CSE_MICROGATEWAY_PRO_SMALL_1: EngineSimpleInfoFlavor{
			value: "cse.microgateway.pro.small.1",
		},
		CSE_MICROGATEWAY_PRO_MEDIUM_1: EngineSimpleInfoFlavor{
			value: "cse.microgateway.pro.medium.1",
		},
		CSE_MICROGATEWAY_PRO_LARGE_1: EngineSimpleInfoFlavor{
			value: "cse.microgateway.pro.large.1",
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
