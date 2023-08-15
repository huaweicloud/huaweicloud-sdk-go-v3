package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OpenApiParaForPublish struct {

	// api编号
	ApiId *string `json:"api_id,omitempty"`

	// 操作类型, 包括发布/下线/停用/恢复
	Action *OpenApiParaForPublishAction `json:"action,omitempty"`

	// 截止时间。仅定期执行需要此参数，默认服务器当前时间三天后。
	Time *string `json:"time,omitempty"`
}

func (o OpenApiParaForPublish) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenApiParaForPublish struct{}"
	}

	return strings.Join([]string{"OpenApiParaForPublish", string(data)}, " ")
}

type OpenApiParaForPublishAction struct {
	value string
}

type OpenApiParaForPublishActionEnum struct {
	PUBLISH   OpenApiParaForPublishAction
	UNPUBLISH OpenApiParaForPublishAction
	STOP      OpenApiParaForPublishAction
	RECOVER   OpenApiParaForPublishAction
}

func GetOpenApiParaForPublishActionEnum() OpenApiParaForPublishActionEnum {
	return OpenApiParaForPublishActionEnum{
		PUBLISH: OpenApiParaForPublishAction{
			value: "PUBLISH",
		},
		UNPUBLISH: OpenApiParaForPublishAction{
			value: "UNPUBLISH",
		},
		STOP: OpenApiParaForPublishAction{
			value: "STOP",
		},
		RECOVER: OpenApiParaForPublishAction{
			value: "RECOVER",
		},
	}
}

func (c OpenApiParaForPublishAction) Value() string {
	return c.value
}

func (c OpenApiParaForPublishAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenApiParaForPublishAction) UnmarshalJSON(b []byte) error {
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
