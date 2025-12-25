package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentDetailInfo 插件定义简介
type ComponentDetailInfo struct {

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
	Upgrade *ComponentDetailInfoUpgrade `json:"upgrade,omitempty"`

	// 更新失败的消息
	UpgradeFailMessage *string `json:"upgrade_fail_message,omitempty"`

	// 安全云脑版本
	Version *string `json:"version,omitempty"`
}

func (o ComponentDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentDetailInfo struct{}"
	}

	return strings.Join([]string{"ComponentDetailInfo", string(data)}, " ")
}

type ComponentDetailInfoUpgrade struct {
	value string
}

type ComponentDetailInfoUpgradeEnum struct {
	NEED_UPGRADE    ComponentDetailInfoUpgrade
	UPGRADING       ComponentDetailInfoUpgrade
	SUCCESS_UPGRADE ComponentDetailInfoUpgrade
	FAIL_UPGRADE    ComponentDetailInfoUpgrade
}

func GetComponentDetailInfoUpgradeEnum() ComponentDetailInfoUpgradeEnum {
	return ComponentDetailInfoUpgradeEnum{
		NEED_UPGRADE: ComponentDetailInfoUpgrade{
			value: "NEED_UPGRADE",
		},
		UPGRADING: ComponentDetailInfoUpgrade{
			value: "UPGRADING",
		},
		SUCCESS_UPGRADE: ComponentDetailInfoUpgrade{
			value: "SUCCESS_UPGRADE",
		},
		FAIL_UPGRADE: ComponentDetailInfoUpgrade{
			value: "FAIL_UPGRADE",
		},
	}
}

func (c ComponentDetailInfoUpgrade) Value() string {
	return c.value
}

func (c ComponentDetailInfoUpgrade) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentDetailInfoUpgrade) UnmarshalJSON(b []byte) error {
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
