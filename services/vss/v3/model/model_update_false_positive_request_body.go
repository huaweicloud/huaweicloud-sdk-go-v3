package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFalsePositiveRequestBody struct {

	// 漏洞ID
	VulnId string `json:"vuln_id"`

	// 误报确认人
	Provider *string `json:"provider,omitempty"`

	// 误报确认理由
	Reason *string `json:"reason,omitempty"`

	// 对漏洞的操作:   * false_report - 更新漏洞状态为误报，并忽略   * repairing - 更新漏洞状态未修复
	VulnStatus *UpdateFalsePositiveRequestBodyVulnStatus `json:"vuln_status,omitempty"`
}

func (o UpdateFalsePositiveRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFalsePositiveRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFalsePositiveRequestBody", string(data)}, " ")
}

type UpdateFalsePositiveRequestBodyVulnStatus struct {
	value string
}

type UpdateFalsePositiveRequestBodyVulnStatusEnum struct {
	FALSE_REPORT UpdateFalsePositiveRequestBodyVulnStatus
	REPAIRING    UpdateFalsePositiveRequestBodyVulnStatus
}

func GetUpdateFalsePositiveRequestBodyVulnStatusEnum() UpdateFalsePositiveRequestBodyVulnStatusEnum {
	return UpdateFalsePositiveRequestBodyVulnStatusEnum{
		FALSE_REPORT: UpdateFalsePositiveRequestBodyVulnStatus{
			value: "false_report",
		},
		REPAIRING: UpdateFalsePositiveRequestBodyVulnStatus{
			value: "repairing",
		},
	}
}

func (c UpdateFalsePositiveRequestBodyVulnStatus) Value() string {
	return c.value
}

func (c UpdateFalsePositiveRequestBodyVulnStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFalsePositiveRequestBodyVulnStatus) UnmarshalJSON(b []byte) error {
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
