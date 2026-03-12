package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DeleteThumbnailResult struct {

	// 删除截图结果对应的媒资id
	AssetId *string `json:"asset_id,omitempty"`

	// 删除截图结果对应的任务id
	TaskId *string `json:"task_id,omitempty"`

	// 删除状态。  取值如下： - DELETED：已删除。 - FAILED：删除失败。 - UNKNOW：未知原因
	Status *DeleteThumbnailResultStatus `json:"status,omitempty"`
}

func (o DeleteThumbnailResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteThumbnailResult struct{}"
	}

	return strings.Join([]string{"DeleteThumbnailResult", string(data)}, " ")
}

type DeleteThumbnailResultStatus struct {
	value string
}

type DeleteThumbnailResultStatusEnum struct {
	FAILED  DeleteThumbnailResultStatus
	DELETED DeleteThumbnailResultStatus
	UNKNOW  DeleteThumbnailResultStatus
}

func GetDeleteThumbnailResultStatusEnum() DeleteThumbnailResultStatusEnum {
	return DeleteThumbnailResultStatusEnum{
		FAILED: DeleteThumbnailResultStatus{
			value: "FAILED",
		},
		DELETED: DeleteThumbnailResultStatus{
			value: "DELETED",
		},
		UNKNOW: DeleteThumbnailResultStatus{
			value: "UNKNOW",
		},
	}
}

func (c DeleteThumbnailResultStatus) Value() string {
	return c.value
}

func (c DeleteThumbnailResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteThumbnailResultStatus) UnmarshalJSON(b []byte) error {
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
