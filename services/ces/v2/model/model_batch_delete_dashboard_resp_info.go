package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchDeleteDashboardRespInfo struct {

	// **参数描述**： 监控看板id **约束限制** 不涉及 **取值范围** 以db开头，包含22个字母和数字，长度为24个字符 **默认取值** 不涉及
	DashboardId *string `json:"dashboard_id,omitempty"`

	// 处理结果, successful: 成功, error: 失败
	RetStatus *BatchDeleteDashboardRespInfoRetStatus `json:"ret_status,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BatchDeleteDashboardRespInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDashboardRespInfo struct{}"
	}

	return strings.Join([]string{"BatchDeleteDashboardRespInfo", string(data)}, " ")
}

type BatchDeleteDashboardRespInfoRetStatus struct {
	value string
}

type BatchDeleteDashboardRespInfoRetStatusEnum struct {
	SUCCESSFUL BatchDeleteDashboardRespInfoRetStatus
	ERROR      BatchDeleteDashboardRespInfoRetStatus
}

func GetBatchDeleteDashboardRespInfoRetStatusEnum() BatchDeleteDashboardRespInfoRetStatusEnum {
	return BatchDeleteDashboardRespInfoRetStatusEnum{
		SUCCESSFUL: BatchDeleteDashboardRespInfoRetStatus{
			value: "successful",
		},
		ERROR: BatchDeleteDashboardRespInfoRetStatus{
			value: "error",
		},
	}
}

func (c BatchDeleteDashboardRespInfoRetStatus) Value() string {
	return c.value
}

func (c BatchDeleteDashboardRespInfoRetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteDashboardRespInfoRetStatus) UnmarshalJSON(b []byte) error {
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
