package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPlatformManagedResponse Response Object
type ShowPlatformManagedResponse struct {

	// 创建时间.
	CreateTime *int64 `json:"create_time,omitempty"`

	// 区域.
	DwRegion *string `json:"dw_region,omitempty"`

	// 平台租户ID.
	PlatformManagedDomainId *string `json:"platform_managed_domain_id,omitempty"`

	// 发布状态.
	PublishStatus *ShowPlatformManagedResponsePublishStatus `json:"publish_status,omitempty"`

	// 业务租户ID.
	TenantManagedDomainId *string `json:"tenant_managed_domain_id,omitempty"`

	// 更新时间.
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 是否在白名单中.
	Whitelist      *bool `json:"whitelist,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowPlatformManagedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlatformManagedResponse struct{}"
	}

	return strings.Join([]string{"ShowPlatformManagedResponse", string(data)}, " ")
}

type ShowPlatformManagedResponsePublishStatus struct {
	value string
}

type ShowPlatformManagedResponsePublishStatusEnum struct {
	DISABLE ShowPlatformManagedResponsePublishStatus
	PROD    ShowPlatformManagedResponsePublishStatus
	GRAY    ShowPlatformManagedResponsePublishStatus
}

func GetShowPlatformManagedResponsePublishStatusEnum() ShowPlatformManagedResponsePublishStatusEnum {
	return ShowPlatformManagedResponsePublishStatusEnum{
		DISABLE: ShowPlatformManagedResponsePublishStatus{
			value: "DISABLE",
		},
		PROD: ShowPlatformManagedResponsePublishStatus{
			value: "PROD",
		},
		GRAY: ShowPlatformManagedResponsePublishStatus{
			value: "GRAY",
		},
	}
}

func (c ShowPlatformManagedResponsePublishStatus) Value() string {
	return c.value
}

func (c ShowPlatformManagedResponsePublishStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPlatformManagedResponsePublishStatus) UnmarshalJSON(b []byte) error {
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
