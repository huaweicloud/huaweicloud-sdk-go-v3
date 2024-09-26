package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DisassociateGlobalConnectionBandwidthInstanceResponseInfo struct {

	// 功能说明：实例ID。 取值范围：1-36个字符，支持数字、字母、_(下划线)、-（中划线）
	ResourceId string `json:"resource_id"`

	// 功能说明：实例类型。
	ResourceType string `json:"resource_type"`

	// 功能说明：实例所在region，不填默认\"global\"。
	RegionId *string `json:"region_id,omitempty"`

	// 功能说明：实例所在region对应的projectId。
	ProjectId string `json:"project_id"`

	// 功能说明：解绑操作成功还是失败。 - success: 成功 - fail: 失败
	Result *DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult `json:"result,omitempty"`

	// 功能说明：解绑操作如果失败，具体的错误信息。
	Message *string `json:"message,omitempty"`
}

func (o DisassociateGlobalConnectionBandwidthInstanceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateGlobalConnectionBandwidthInstanceResponseInfo struct{}"
	}

	return strings.Join([]string{"DisassociateGlobalConnectionBandwidthInstanceResponseInfo", string(data)}, " ")
}

type DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult struct {
	value string
}

type DisassociateGlobalConnectionBandwidthInstanceResponseInfoResultEnum struct {
	SUCCESS DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult
	FAIL    DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult
}

func GetDisassociateGlobalConnectionBandwidthInstanceResponseInfoResultEnum() DisassociateGlobalConnectionBandwidthInstanceResponseInfoResultEnum {
	return DisassociateGlobalConnectionBandwidthInstanceResponseInfoResultEnum{
		SUCCESS: DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult{
			value: "success",
		},
		FAIL: DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult{
			value: "fail",
		},
	}
}

func (c DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult) Value() string {
	return c.value
}

func (c DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisassociateGlobalConnectionBandwidthInstanceResponseInfoResult) UnmarshalJSON(b []byte) error {
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
