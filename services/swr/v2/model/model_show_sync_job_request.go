package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ShowSyncJobRequest struct {
	ContentType ShowSyncJobRequestContentType `json:"Content-Type"`
	Namespace   string                        `json:"namespace"`
	Repository  string                        `json:"repository"`
	Filter      string                        `json:"filter"`
}

func (o ShowSyncJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSyncJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSyncJobRequest", string(data)}, " ")
}

type ShowSyncJobRequestContentType struct {
	value string
}

type ShowSyncJobRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowSyncJobRequestContentType
	APPLICATION_JSON             ShowSyncJobRequestContentType
}

func GetShowSyncJobRequestContentTypeEnum() ShowSyncJobRequestContentTypeEnum {
	return ShowSyncJobRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowSyncJobRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowSyncJobRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowSyncJobRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowSyncJobRequestContentType) UnmarshalJSON(b []byte) error {
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
