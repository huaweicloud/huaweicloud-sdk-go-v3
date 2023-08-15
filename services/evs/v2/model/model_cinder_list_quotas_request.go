package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CinderListQuotasRequest Request Object
type CinderListQuotasRequest struct {

	// 目标的项目ID。与project_id保持一致即可。
	TargetProjectId string `json:"target_project_id"`

	// 是否查询配额详细信息。当前只支持传true。
	Usage CinderListQuotasRequestUsage `json:"usage"`
}

func (o CinderListQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListQuotasRequest struct{}"
	}

	return strings.Join([]string{"CinderListQuotasRequest", string(data)}, " ")
}

type CinderListQuotasRequestUsage struct {
	value string
}

type CinderListQuotasRequestUsageEnum struct {
	TRUE CinderListQuotasRequestUsage
}

func GetCinderListQuotasRequestUsageEnum() CinderListQuotasRequestUsageEnum {
	return CinderListQuotasRequestUsageEnum{
		TRUE: CinderListQuotasRequestUsage{
			value: "True",
		},
	}
}

func (c CinderListQuotasRequestUsage) Value() string {
	return c.value
}

func (c CinderListQuotasRequestUsage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CinderListQuotasRequestUsage) UnmarshalJSON(b []byte) error {
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
