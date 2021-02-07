package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type UpdateRepoRequest struct {
	ContentType UpdateRepoRequestContentType `json:"Content-Type"`
	Namespace   string                       `json:"namespace"`
	Repository  string                       `json:"repository"`
	Body        *UpdateRepoRequestBody       `json:"body,omitempty"`
}

func (o UpdateRepoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRepoRequest struct{}"
	}

	return strings.Join([]string{"UpdateRepoRequest", string(data)}, " ")
}

type UpdateRepoRequestContentType struct {
	value string
}

type UpdateRepoRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 UpdateRepoRequestContentType
	APPLICATION_JSON             UpdateRepoRequestContentType
}

func GetUpdateRepoRequestContentTypeEnum() UpdateRepoRequestContentTypeEnum {
	return UpdateRepoRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: UpdateRepoRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: UpdateRepoRequestContentType{
			value: "application/json",
		},
	}
}

func (c UpdateRepoRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateRepoRequestContentType) UnmarshalJSON(b []byte) error {
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
