package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UnreasonablePermission 权限控制合理性检测结果。
type UnreasonablePermission struct {

	// 检测结果 * NO_RISK 无风险 * MEDIUM_RISK 中风险 * HIGH_RISK 高风险 * NOT_SCANNED 未扫描
	Result *UnreasonablePermissionResult `json:"result,omitempty"`

	// 存在风险的权限控制数量。
	Count *int32 `json:"count,omitempty"`
}

func (o UnreasonablePermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnreasonablePermission struct{}"
	}

	return strings.Join([]string{"UnreasonablePermission", string(data)}, " ")
}

type UnreasonablePermissionResult struct {
	value string
}

type UnreasonablePermissionResultEnum struct {
	NO_RISK     UnreasonablePermissionResult
	MEDIUM_RISK UnreasonablePermissionResult
	HIGH_RISK   UnreasonablePermissionResult
	NOT_SCANNED UnreasonablePermissionResult
}

func GetUnreasonablePermissionResultEnum() UnreasonablePermissionResultEnum {
	return UnreasonablePermissionResultEnum{
		NO_RISK: UnreasonablePermissionResult{
			value: "NO_RISK",
		},
		MEDIUM_RISK: UnreasonablePermissionResult{
			value: "MEDIUM_RISK",
		},
		HIGH_RISK: UnreasonablePermissionResult{
			value: "HIGH_RISK",
		},
		NOT_SCANNED: UnreasonablePermissionResult{
			value: "NOT_SCANNED",
		},
	}
}

func (c UnreasonablePermissionResult) Value() string {
	return c.value
}

func (c UnreasonablePermissionResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UnreasonablePermissionResult) UnmarshalJSON(b []byte) error {
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
