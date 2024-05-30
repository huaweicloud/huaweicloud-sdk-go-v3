package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DsProcessImportResultVo struct {

	// 标识本次导入的唯一值，用于查询导入结果。 枚举值：   - importing: 导入中   - fail: 导入失败   - success: 导入成功
	Status *DsProcessImportResultVoStatus `json:"status,omitempty"`

	Group *BatchOperationVo `json:"group,omitempty"`

	// 当前进度。
	Rate *string `json:"rate,omitempty"`
}

func (o DsProcessImportResultVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DsProcessImportResultVo struct{}"
	}

	return strings.Join([]string{"DsProcessImportResultVo", string(data)}, " ")
}

type DsProcessImportResultVoStatus struct {
	value string
}

type DsProcessImportResultVoStatusEnum struct {
	IMPORTING DsProcessImportResultVoStatus
	FAIL      DsProcessImportResultVoStatus
	SUCCESS   DsProcessImportResultVoStatus
}

func GetDsProcessImportResultVoStatusEnum() DsProcessImportResultVoStatusEnum {
	return DsProcessImportResultVoStatusEnum{
		IMPORTING: DsProcessImportResultVoStatus{
			value: "importing",
		},
		FAIL: DsProcessImportResultVoStatus{
			value: "fail",
		},
		SUCCESS: DsProcessImportResultVoStatus{
			value: "success",
		},
	}
}

func (c DsProcessImportResultVoStatus) Value() string {
	return c.value
}

func (c DsProcessImportResultVoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DsProcessImportResultVoStatus) UnmarshalJSON(b []byte) error {
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
