package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchAddPermissionRequest 批量添加终端节点服务白名单列表。
type BatchAddPermissionRequest struct {

	// 终端节点服务白名单列表
	Permissions []EpsAddPermissionRequest `json:"permissions"`

	// 终端节点服务白名单类型。 ● domainId：基于账户ID配置终端节点服务白名单。 ● orgPath：基于账户所在组织路径配置终端节点服务白名单。
	PermissionType *BatchAddPermissionRequestPermissionType `json:"permission_type,omitempty"`
}

func (o BatchAddPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddPermissionRequest struct{}"
	}

	return strings.Join([]string{"BatchAddPermissionRequest", string(data)}, " ")
}

type BatchAddPermissionRequestPermissionType struct {
	value string
}

type BatchAddPermissionRequestPermissionTypeEnum struct {
	DOMAIN_ID BatchAddPermissionRequestPermissionType
	ORG_PATH  BatchAddPermissionRequestPermissionType
}

func GetBatchAddPermissionRequestPermissionTypeEnum() BatchAddPermissionRequestPermissionTypeEnum {
	return BatchAddPermissionRequestPermissionTypeEnum{
		DOMAIN_ID: BatchAddPermissionRequestPermissionType{
			value: "domainId",
		},
		ORG_PATH: BatchAddPermissionRequestPermissionType{
			value: "orgPath",
		},
	}
}

func (c BatchAddPermissionRequestPermissionType) Value() string {
	return c.value
}

func (c BatchAddPermissionRequestPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddPermissionRequestPermissionType) UnmarshalJSON(b []byte) error {
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
