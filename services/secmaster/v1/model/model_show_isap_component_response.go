package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowIsapComponentResponse Response Object
type ShowIsapComponentResponse struct {

	// 组件id.
	ComponentId *string `json:"component_id,omitempty"`

	// 组件名称
	ComponentName *string `json:"component_name,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 相关描述
	Description *string `json:"description,omitempty"`

	// 历史版本
	HistoryVersion *string `json:"history_version,omitempty"`

	// 保持
	Maintainer *string `json:"maintainer,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 升级 - NEED_UPGRADE 需要升级 - UPGRADING 升级中 - SUCCESS_UPGRADE 升级成功 - FAIL_UPGRADE 升级失败  **约束限制** 不涉及 **取值范围**: - NEED_UPGRADE - UPGRADING - SUCCESS_UPGRADE - FAIL_UPGRADE  **默认值** 不涉及
	Upgrade *ShowIsapComponentResponseUpgrade `json:"upgrade,omitempty"`

	// 更新失败的消息
	UpgradeFailMessage *string `json:"upgrade_fail_message,omitempty"`

	// 安全云脑版本
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIsapComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIsapComponentResponse struct{}"
	}

	return strings.Join([]string{"ShowIsapComponentResponse", string(data)}, " ")
}

type ShowIsapComponentResponseUpgrade struct {
	value string
}

type ShowIsapComponentResponseUpgradeEnum struct {
	NEED_UPGRADE    ShowIsapComponentResponseUpgrade
	UPGRADING       ShowIsapComponentResponseUpgrade
	SUCCESS_UPGRADE ShowIsapComponentResponseUpgrade
	FAIL_UPGRADE    ShowIsapComponentResponseUpgrade
}

func GetShowIsapComponentResponseUpgradeEnum() ShowIsapComponentResponseUpgradeEnum {
	return ShowIsapComponentResponseUpgradeEnum{
		NEED_UPGRADE: ShowIsapComponentResponseUpgrade{
			value: "NEED_UPGRADE",
		},
		UPGRADING: ShowIsapComponentResponseUpgrade{
			value: "UPGRADING",
		},
		SUCCESS_UPGRADE: ShowIsapComponentResponseUpgrade{
			value: "SUCCESS_UPGRADE",
		},
		FAIL_UPGRADE: ShowIsapComponentResponseUpgrade{
			value: "FAIL_UPGRADE",
		},
	}
}

func (c ShowIsapComponentResponseUpgrade) Value() string {
	return c.value
}

func (c ShowIsapComponentResponseUpgrade) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowIsapComponentResponseUpgrade) UnmarshalJSON(b []byte) error {
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
