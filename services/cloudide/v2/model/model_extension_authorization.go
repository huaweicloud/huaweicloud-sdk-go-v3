package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExtensionAuthorization struct {

	// 插件版本
	ExtensionVersion string `json:"extension_version"`

	// 插件标识(发布者.插件名)
	Identifier string `json:"identifier"`

	// CodeArtsIDEOnline实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 插件状态。 - AGREE 同意 - REJECT 不同意 - UNKNOWN 未知（下次重新询问）
	Status ExtensionAuthorizationStatus `json:"status"`
}

func (o ExtensionAuthorization) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionAuthorization struct{}"
	}

	return strings.Join([]string{"ExtensionAuthorization", string(data)}, " ")
}

type ExtensionAuthorizationStatus struct {
	value string
}

type ExtensionAuthorizationStatusEnum struct {
	AGREE   ExtensionAuthorizationStatus
	REJECT  ExtensionAuthorizationStatus
	UNKNOWN ExtensionAuthorizationStatus
}

func GetExtensionAuthorizationStatusEnum() ExtensionAuthorizationStatusEnum {
	return ExtensionAuthorizationStatusEnum{
		AGREE: ExtensionAuthorizationStatus{
			value: "AGREE",
		},
		REJECT: ExtensionAuthorizationStatus{
			value: "REJECT",
		},
		UNKNOWN: ExtensionAuthorizationStatus{
			value: "UNKNOWN",
		},
	}
}

func (c ExtensionAuthorizationStatus) Value() string {
	return c.value
}

func (c ExtensionAuthorizationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtensionAuthorizationStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
