package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdateWidgetInfo **参数解释** 批量更新监控视图结果
type BatchUpdateWidgetInfo struct {

	// **参数解释** 视图id **取值范围** 字符串必须以wg开头，后跟22个字母和数字，总长度为24个字符
	WidgetId *string `json:"widget_id,omitempty"`

	// **参数解释** 修改结果 **取值范围** 枚举值： - successful 成功 - error 失败
	RetStatus *BatchUpdateWidgetInfoRetStatus `json:"ret_status,omitempty"`

	// **参数解释** 如果失败则返回失败信息 **取值范围** 长度为[1,2048]个字符
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BatchUpdateWidgetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateWidgetInfo struct{}"
	}

	return strings.Join([]string{"BatchUpdateWidgetInfo", string(data)}, " ")
}

type BatchUpdateWidgetInfoRetStatus struct {
	value string
}

type BatchUpdateWidgetInfoRetStatusEnum struct {
	SUCCESSFUL BatchUpdateWidgetInfoRetStatus
	ERROR      BatchUpdateWidgetInfoRetStatus
}

func GetBatchUpdateWidgetInfoRetStatusEnum() BatchUpdateWidgetInfoRetStatusEnum {
	return BatchUpdateWidgetInfoRetStatusEnum{
		SUCCESSFUL: BatchUpdateWidgetInfoRetStatus{
			value: "successful",
		},
		ERROR: BatchUpdateWidgetInfoRetStatus{
			value: "error",
		},
	}
}

func (c BatchUpdateWidgetInfoRetStatus) Value() string {
	return c.value
}

func (c BatchUpdateWidgetInfoRetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateWidgetInfoRetStatus) UnmarshalJSON(b []byte) error {
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
