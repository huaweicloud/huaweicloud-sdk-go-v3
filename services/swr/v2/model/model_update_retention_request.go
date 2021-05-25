package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateRetentionRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType UpdateRetentionRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
	// 镜像老化规则id

	RetentionId int32 `json:"retention_id"`

	Body *UpdateRetentionRequestBody `json:"body,omitempty"`
}

func (o UpdateRetentionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRetentionRequest struct{}"
	}

	return strings.Join([]string{"UpdateRetentionRequest", string(data)}, " ")
}

type UpdateRetentionRequestContentType struct {
	value string
}

type UpdateRetentionRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 UpdateRetentionRequestContentType
	APPLICATION_JSON             UpdateRetentionRequestContentType
}

func GetUpdateRetentionRequestContentTypeEnum() UpdateRetentionRequestContentTypeEnum {
	return UpdateRetentionRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: UpdateRetentionRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: UpdateRetentionRequestContentType{
			value: "application/json",
		},
	}
}

func (c UpdateRetentionRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateRetentionRequestContentType) UnmarshalJSON(b []byte) error {
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
