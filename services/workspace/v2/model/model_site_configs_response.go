package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SiteConfigsResponse 站点配置
type SiteConfigsResponse struct {

	// 站点id
	SiteId *string `json:"site_id,omitempty"`

	// 配置状态。 - CENTER： 中心初始化 - IES： 边缘初始化
	SiteType *SiteConfigsResponseSiteType `json:"site_type,omitempty"`

	// 站点名称
	SiteName *string `json:"site_name,omitempty"`

	// 云办公服务的状态。 - PREPARING：准备初始化服务 - SUBSCRIBING：初始化服务中 - SUBSCRIBED：已初始化服务 - SUBSCRIPTION_FAILED：初始化服务失败 - DEREGISTERING：清理资源中 - DEREGISTRATION_FAILED：清理资源失败 - RECYCLING：清理资源中。 - RECYCLED：清理资源成功。 - RECYCLE_FAILED：清理资源失败。 - CLOSED：已销户未初始化服务
	Status *SiteConfigsResponseStatus `json:"status,omitempty"`

	// 互联网和专线切换任务的状态。 - init： 初始化 - 开通服务后的初始状态 - available： 可用 - 执行过任务且成功后恢复的正常状态 - internetOpening： 开启中 - 开通互联网接入开启中 - dedicatedOpening： 开启中 - 开通专线接入开启中 - internetOpenFailed： 开启失败 - 开通互联网接入开启失败 - dedicatedOpenFailed： 开启失败 - 开通专线接入开启失败 - openSuccess： 开启成功 - 开通互联网接入成功 - internetClosing： 关闭中 - 关闭互联网接入关闭中 - dedicatedClosing： 关闭中 - 关闭专线接入关闭中 - internetCloseFailed： 关闭失败 - 关闭互联网接入方式失败 - dedicatedCloseFailed： 关闭失败 - 关闭专线接入方式失败 - closeSuccess： 关闭成功 - 关闭接入方式成功 - internetAccessPortModifying： 互联网接入端口修改中 - internetAccessPortModifyFailed: 端口修改失败
	AccessStatus *string `json:"access_status,omitempty"`

	// 配置状态。 - \"0\"： 开通服务成功，且对接AD成功 - \"1\"： 开通服务成功，但AD配置失败 - \"2\"： 开通服务成功，但AD配置失败后存在其他错误 - \"3\"： 开通服务成功，AD未开启对接
	ConfigStatus *string `json:"config_status,omitempty"`

	InfrastructureSecurityGroup *SecurityGroup `json:"infrastructure_security_group,omitempty"`

	DesktopSecurityGroup *SecurityGroup `json:"desktop_security_group,omitempty"`

	// 开通服务资源使用的可用分区
	AvailabilityZones *[]string `json:"availability_zones,omitempty"`

	// 开通服务或取消服务的任务ID
	JobId *string `json:"job_id,omitempty"`

	// 初始化服务或清理资源的进度，格式为100%
	Progress *string `json:"progress,omitempty"`

	// 失败错误码
	FailCode *int32 `json:"fail_code,omitempty"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	NetworkConfig *NetworkConfig `json:"network_config,omitempty"`

	AccessConfig *AccessConfig `json:"access_config,omitempty"`

	// 是否可以取消服务。
	Closable *bool `json:"closable,omitempty"`
}

func (o SiteConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteConfigsResponse struct{}"
	}

	return strings.Join([]string{"SiteConfigsResponse", string(data)}, " ")
}

type SiteConfigsResponseSiteType struct {
	value string
}

type SiteConfigsResponseSiteTypeEnum struct {
	CENTER SiteConfigsResponseSiteType
	IES    SiteConfigsResponseSiteType
}

func GetSiteConfigsResponseSiteTypeEnum() SiteConfigsResponseSiteTypeEnum {
	return SiteConfigsResponseSiteTypeEnum{
		CENTER: SiteConfigsResponseSiteType{
			value: "CENTER",
		},
		IES: SiteConfigsResponseSiteType{
			value: "IES",
		},
	}
}

func (c SiteConfigsResponseSiteType) Value() string {
	return c.value
}

func (c SiteConfigsResponseSiteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteConfigsResponseSiteType) UnmarshalJSON(b []byte) error {
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

type SiteConfigsResponseStatus struct {
	value string
}

type SiteConfigsResponseStatusEnum struct {
	PREPARING             SiteConfigsResponseStatus
	SUBSCRIBING           SiteConfigsResponseStatus
	SUBSCRIBED            SiteConfigsResponseStatus
	SUBSCRIPTION_FAILED   SiteConfigsResponseStatus
	DEREGISTERING         SiteConfigsResponseStatus
	DEREGISTRATION_FAILED SiteConfigsResponseStatus
	RECYCLING             SiteConfigsResponseStatus
	RECYCLED              SiteConfigsResponseStatus
	RECYCLE_FAILED        SiteConfigsResponseStatus
	CLOSED                SiteConfigsResponseStatus
}

func GetSiteConfigsResponseStatusEnum() SiteConfigsResponseStatusEnum {
	return SiteConfigsResponseStatusEnum{
		PREPARING: SiteConfigsResponseStatus{
			value: "PREPARING",
		},
		SUBSCRIBING: SiteConfigsResponseStatus{
			value: "SUBSCRIBING",
		},
		SUBSCRIBED: SiteConfigsResponseStatus{
			value: "SUBSCRIBED",
		},
		SUBSCRIPTION_FAILED: SiteConfigsResponseStatus{
			value: "SUBSCRIPTION_FAILED",
		},
		DEREGISTERING: SiteConfigsResponseStatus{
			value: "DEREGISTERING",
		},
		DEREGISTRATION_FAILED: SiteConfigsResponseStatus{
			value: "DEREGISTRATION_FAILED",
		},
		RECYCLING: SiteConfigsResponseStatus{
			value: "RECYCLING",
		},
		RECYCLED: SiteConfigsResponseStatus{
			value: "RECYCLED",
		},
		RECYCLE_FAILED: SiteConfigsResponseStatus{
			value: "RECYCLE_FAILED",
		},
		CLOSED: SiteConfigsResponseStatus{
			value: "CLOSED",
		},
	}
}

func (c SiteConfigsResponseStatus) Value() string {
	return c.value
}

func (c SiteConfigsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteConfigsResponseStatus) UnmarshalJSON(b []byte) error {
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
