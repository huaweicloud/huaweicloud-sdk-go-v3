package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DsExportResultVoData 导出结果对象。
type DsExportResultVoData struct {

	// 标识本次导出的唯一值，用于查询导入结果。importing(导出中)、fail(导出失败)、success(导出成功)。 枚举值：   - importing: 导出中   - fail: 导出失败   - success: 导出成功
	Status *DsExportResultVoDataStatus `json:"status,omitempty"`

	Group *BatchOperationVo `json:"group,omitempty"`

	// 当前进度。
	Rate *string `json:"rate,omitempty"`
}

func (o DsExportResultVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DsExportResultVoData struct{}"
	}

	return strings.Join([]string{"DsExportResultVoData", string(data)}, " ")
}

type DsExportResultVoDataStatus struct {
	value string
}

type DsExportResultVoDataStatusEnum struct {
	IMPORTING DsExportResultVoDataStatus
	FAIL      DsExportResultVoDataStatus
	SUCCESS   DsExportResultVoDataStatus
}

func GetDsExportResultVoDataStatusEnum() DsExportResultVoDataStatusEnum {
	return DsExportResultVoDataStatusEnum{
		IMPORTING: DsExportResultVoDataStatus{
			value: "importing",
		},
		FAIL: DsExportResultVoDataStatus{
			value: "fail",
		},
		SUCCESS: DsExportResultVoDataStatus{
			value: "success",
		},
	}
}

func (c DsExportResultVoDataStatus) Value() string {
	return c.value
}

func (c DsExportResultVoDataStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DsExportResultVoDataStatus) UnmarshalJSON(b []byte) error {
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
