package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpgradeActionInfo 升级操作信息。
type UpgradeActionInfo struct {

	// 升级操作,upgrade=升级,upgradeAutoCommit=升级自动提交,commit=提交,rollback=回滚。
	UpgradeAction *UpgradeActionInfoUpgradeAction `json:"upgrade_action,omitempty"`

	// 可用，不可用。
	Enable *bool `json:"enable,omitempty"`
}

func (o UpgradeActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeActionInfo struct{}"
	}

	return strings.Join([]string{"UpgradeActionInfo", string(data)}, " ")
}

type UpgradeActionInfoUpgradeAction struct {
	value string
}

type UpgradeActionInfoUpgradeActionEnum struct {
	UPGRADE             UpgradeActionInfoUpgradeAction
	UPGRADE_AUTO_COMMIT UpgradeActionInfoUpgradeAction
	COMMIT              UpgradeActionInfoUpgradeAction
	ROLLBACK            UpgradeActionInfoUpgradeAction
}

func GetUpgradeActionInfoUpgradeActionEnum() UpgradeActionInfoUpgradeActionEnum {
	return UpgradeActionInfoUpgradeActionEnum{
		UPGRADE: UpgradeActionInfoUpgradeAction{
			value: "upgrade",
		},
		UPGRADE_AUTO_COMMIT: UpgradeActionInfoUpgradeAction{
			value: "upgradeAutoCommit",
		},
		COMMIT: UpgradeActionInfoUpgradeAction{
			value: "commit",
		},
		ROLLBACK: UpgradeActionInfoUpgradeAction{
			value: "rollback",
		},
	}
}

func (c UpgradeActionInfoUpgradeAction) Value() string {
	return c.value
}

func (c UpgradeActionInfoUpgradeAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradeActionInfoUpgradeAction) UnmarshalJSON(b []byte) error {
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
