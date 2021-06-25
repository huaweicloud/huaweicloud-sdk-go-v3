package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DeleteResult struct {
	// 媒体ID<br/>

	AssetId *string `json:"asset_id,omitempty"`
	// 删除状态<br/>

	Status *DeleteResultStatus `json:"status,omitempty"`
}

func (o DeleteResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteResult struct{}"
	}

	return strings.Join([]string{"DeleteResult", string(data)}, " ")
}

type DeleteResultStatus struct {
	value string
}

type DeleteResultStatusEnum struct {
	FAILED  DeleteResultStatus
	DELETED DeleteResultStatus
	UNKNOW  DeleteResultStatus
}

func GetDeleteResultStatusEnum() DeleteResultStatusEnum {
	return DeleteResultStatusEnum{
		FAILED: DeleteResultStatus{
			value: "FAILED",
		},
		DELETED: DeleteResultStatus{
			value: "DELETED",
		},
		UNKNOW: DeleteResultStatus{
			value: "UNKNOW",
		},
	}
}

func (c DeleteResultStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DeleteResultStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
