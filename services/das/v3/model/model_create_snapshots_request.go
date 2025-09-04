package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSnapshotsRequest Request Object
type CreateSnapshotsRequest struct {

	// 连接id
	ConnectionId string `json:"connection_id"`

	// 语言
	XLanguage *CreateSnapshotsRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateSnapshotsRequestBody `json:"body,omitempty"`
}

func (o CreateSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"CreateSnapshotsRequest", string(data)}, " ")
}

type CreateSnapshotsRequestXLanguage struct {
	value string
}

type CreateSnapshotsRequestXLanguageEnum struct {
	ZH_CN CreateSnapshotsRequestXLanguage
	EN_US CreateSnapshotsRequestXLanguage
}

func GetCreateSnapshotsRequestXLanguageEnum() CreateSnapshotsRequestXLanguageEnum {
	return CreateSnapshotsRequestXLanguageEnum{
		ZH_CN: CreateSnapshotsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateSnapshotsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateSnapshotsRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSnapshotsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSnapshotsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
