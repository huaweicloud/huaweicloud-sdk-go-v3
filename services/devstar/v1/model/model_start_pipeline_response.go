package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type StartPipelineResponse struct {
	// 流水线id

	Id *string `json:"id,omitempty"`
	// 流水线操作成功

	Result         *StartPipelineResponseResult `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o StartPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineResponse struct{}"
	}

	return strings.Join([]string{"StartPipelineResponse", string(data)}, " ")
}

type StartPipelineResponseResult struct {
	value string
}

type StartPipelineResponseResultEnum struct {
	SUCCESS StartPipelineResponseResult
	FAILED  StartPipelineResponseResult
}

func GetStartPipelineResponseResultEnum() StartPipelineResponseResultEnum {
	return StartPipelineResponseResultEnum{
		SUCCESS: StartPipelineResponseResult{
			value: "success",
		},
		FAILED: StartPipelineResponseResult{
			value: "failed",
		},
	}
}

func (c StartPipelineResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartPipelineResponseResult) UnmarshalJSON(b []byte) error {
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
