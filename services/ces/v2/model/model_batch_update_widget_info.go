package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdateWidgetInfo 批量更新监控视图结果
type BatchUpdateWidgetInfo struct {

	// 视图id
	WidgetId *string `json:"widget_id,omitempty"`

	// 修改结果；成功: successful, 失败: error
	RetStatus *BatchUpdateWidgetInfoRetStatus `json:"ret_status,omitempty"`

	// 如果失败则返回失败信息
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
