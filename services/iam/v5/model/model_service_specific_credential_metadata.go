package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ServiceSpecificCredentialMetadata 服务专属凭证元数据。
type ServiceSpecificCredentialMetadata struct {

	// 创建日期和时间，ISO 8601格式。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 关联的服务名称。
	ServiceName string `json:"service_name"`

	// 凭证标识符ID。
	CredentialId string `json:"credential_id"`

	// 凭证状态。
	Status ServiceSpecificCredentialMetadataStatus `json:"status"`

	// 与服务专属凭证关联的IAM用户ID。
	UserId string `json:"user_id"`

	// 关联的IAM用户名称。
	UserName string `json:"user_name"`

	// 过期日期和时间，仅指定credential_age_days时返回。
	ExpiresAt *sdktime.SdkTime `json:"expires_at,omitempty"`

	// 凭证的掩码值。
	CredentialMask string `json:"credential_mask"`

	// 凭证描述。
	Description string `json:"description"`
}

func (o ServiceSpecificCredentialMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSpecificCredentialMetadata struct{}"
	}

	return strings.Join([]string{"ServiceSpecificCredentialMetadata", string(data)}, " ")
}

type ServiceSpecificCredentialMetadataStatus struct {
	value string
}

type ServiceSpecificCredentialMetadataStatusEnum struct {
	ACTIVE   ServiceSpecificCredentialMetadataStatus
	INACTIVE ServiceSpecificCredentialMetadataStatus
	EXPIRED  ServiceSpecificCredentialMetadataStatus
}

func GetServiceSpecificCredentialMetadataStatusEnum() ServiceSpecificCredentialMetadataStatusEnum {
	return ServiceSpecificCredentialMetadataStatusEnum{
		ACTIVE: ServiceSpecificCredentialMetadataStatus{
			value: "active",
		},
		INACTIVE: ServiceSpecificCredentialMetadataStatus{
			value: "inactive",
		},
		EXPIRED: ServiceSpecificCredentialMetadataStatus{
			value: "expired",
		},
	}
}

func (c ServiceSpecificCredentialMetadataStatus) Value() string {
	return c.value
}

func (c ServiceSpecificCredentialMetadataStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServiceSpecificCredentialMetadataStatus) UnmarshalJSON(b []byte) error {
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
