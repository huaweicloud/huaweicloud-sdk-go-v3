package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssetActionResult 资产操作结果
type AssetActionResult struct {

	// 处理状态。 * SUCCESS：成功 * FAILED：失败
	RetStatus *AssetActionResultRetStatus `json:"ret_status,omitempty"`

	// 资产ID列表
	AssetIds *[]string `json:"asset_ids,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`
}

func (o AssetActionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetActionResult struct{}"
	}

	return strings.Join([]string{"AssetActionResult", string(data)}, " ")
}

type AssetActionResultRetStatus struct {
	value string
}

type AssetActionResultRetStatusEnum struct {
	SUCCESS AssetActionResultRetStatus
	FAILED  AssetActionResultRetStatus
}

func GetAssetActionResultRetStatusEnum() AssetActionResultRetStatusEnum {
	return AssetActionResultRetStatusEnum{
		SUCCESS: AssetActionResultRetStatus{
			value: "SUCCESS",
		},
		FAILED: AssetActionResultRetStatus{
			value: "FAILED",
		},
	}
}

func (c AssetActionResultRetStatus) Value() string {
	return c.value
}

func (c AssetActionResultRetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetActionResultRetStatus) UnmarshalJSON(b []byte) error {
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
