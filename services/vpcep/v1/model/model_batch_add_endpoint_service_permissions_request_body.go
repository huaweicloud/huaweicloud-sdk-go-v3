package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchAddEndpointServicePermissionsRequestBody 批量添加终端节点服务白名单列表。
type BatchAddEndpointServicePermissionsRequestBody struct {

	// 终端节点服务白名单列表
	Permissions []EpsAddPermissionRequest `json:"permissions"`

	// 终端节点服务白名单类型。  - domainId：基于账户ID配置终端节点服务白名单。  - orgPath：基于账户所在组织路径配置终端节点服务白名单。
	PermissionType *BatchAddEndpointServicePermissionsRequestBodyPermissionType `json:"permission_type,omitempty"`
}

func (o BatchAddEndpointServicePermissionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddEndpointServicePermissionsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddEndpointServicePermissionsRequestBody", string(data)}, " ")
}

type BatchAddEndpointServicePermissionsRequestBodyPermissionType struct {
	value string
}

type BatchAddEndpointServicePermissionsRequestBodyPermissionTypeEnum struct {
	DOMAIN_ID BatchAddEndpointServicePermissionsRequestBodyPermissionType
	ORG_PATH  BatchAddEndpointServicePermissionsRequestBodyPermissionType
}

func GetBatchAddEndpointServicePermissionsRequestBodyPermissionTypeEnum() BatchAddEndpointServicePermissionsRequestBodyPermissionTypeEnum {
	return BatchAddEndpointServicePermissionsRequestBodyPermissionTypeEnum{
		DOMAIN_ID: BatchAddEndpointServicePermissionsRequestBodyPermissionType{
			value: "domainId",
		},
		ORG_PATH: BatchAddEndpointServicePermissionsRequestBodyPermissionType{
			value: "orgPath",
		},
	}
}

func (c BatchAddEndpointServicePermissionsRequestBodyPermissionType) Value() string {
	return c.value
}

func (c BatchAddEndpointServicePermissionsRequestBodyPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddEndpointServicePermissionsRequestBodyPermissionType) UnmarshalJSON(b []byte) error {
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
