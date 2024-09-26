package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AssociateGlobalConnectionBandwidthInstanceResponseInfo struct {

	// 功能说明：实例ID。 取值范围：1-36个字符，支持数字、字母、_(下划线)、-（中划线）
	ResourceId string `json:"resource_id"`

	// 功能说明：实例类型。
	ResourceType string `json:"resource_type"`

	// 功能说明：实例所在region，不填默认\"global\"。
	RegionId *string `json:"region_id,omitempty"`

	// 功能说明：实例所在region对应的projectId。
	ProjectId string `json:"project_id"`

	// 功能说明：绑定操作成功还是失败。 - success: 成功 - fail: 失败
	Result *AssociateGlobalConnectionBandwidthInstanceResponseInfoResult `json:"result,omitempty"`

	// 功能说明：绑定操作如果失败，具体的错误信息。
	Message *string `json:"message,omitempty"`
}

func (o AssociateGlobalConnectionBandwidthInstanceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGlobalConnectionBandwidthInstanceResponseInfo struct{}"
	}

	return strings.Join([]string{"AssociateGlobalConnectionBandwidthInstanceResponseInfo", string(data)}, " ")
}

type AssociateGlobalConnectionBandwidthInstanceResponseInfoResult struct {
	value string
}

type AssociateGlobalConnectionBandwidthInstanceResponseInfoResultEnum struct {
	SUCCESS AssociateGlobalConnectionBandwidthInstanceResponseInfoResult
	FAIL    AssociateGlobalConnectionBandwidthInstanceResponseInfoResult
}

func GetAssociateGlobalConnectionBandwidthInstanceResponseInfoResultEnum() AssociateGlobalConnectionBandwidthInstanceResponseInfoResultEnum {
	return AssociateGlobalConnectionBandwidthInstanceResponseInfoResultEnum{
		SUCCESS: AssociateGlobalConnectionBandwidthInstanceResponseInfoResult{
			value: "success",
		},
		FAIL: AssociateGlobalConnectionBandwidthInstanceResponseInfoResult{
			value: "fail",
		},
	}
}

func (c AssociateGlobalConnectionBandwidthInstanceResponseInfoResult) Value() string {
	return c.value
}

func (c AssociateGlobalConnectionBandwidthInstanceResponseInfoResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociateGlobalConnectionBandwidthInstanceResponseInfoResult) UnmarshalJSON(b []byte) error {
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
