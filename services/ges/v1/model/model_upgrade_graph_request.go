package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpgradeGraphRequest Request Object
type UpgradeGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 图actionId
	ActionId UpgradeGraphRequestActionId `json:"action_id"`

	Body *UpgradeGraphReq `json:"body,omitempty"`
}

func (o UpgradeGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeGraphRequest struct{}"
	}

	return strings.Join([]string{"UpgradeGraphRequest", string(data)}, " ")
}

type UpgradeGraphRequestActionId struct {
	value string
}

type UpgradeGraphRequestActionIdEnum struct {
	UPGRADE UpgradeGraphRequestActionId
}

func GetUpgradeGraphRequestActionIdEnum() UpgradeGraphRequestActionIdEnum {
	return UpgradeGraphRequestActionIdEnum{
		UPGRADE: UpgradeGraphRequestActionId{
			value: "upgrade",
		},
	}
}

func (c UpgradeGraphRequestActionId) Value() string {
	return c.value
}

func (c UpgradeGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradeGraphRequestActionId) UnmarshalJSON(b []byte) error {
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
