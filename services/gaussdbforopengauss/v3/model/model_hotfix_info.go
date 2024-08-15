package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HotfixInfo 热补丁信息。
type HotfixInfo struct {

	// 热补丁版本。
	Version *string `json:"version,omitempty"`

	// 通用非通用信息,common=通用补丁,certain=定制补丁。
	CommonPatch *HotfixInfoCommonPatch `json:"common_patch,omitempty"`

	// 是否和备份相关。
	BackupSensitive *bool `json:"backup_sensitive,omitempty"`

	// 补丁的描述信息。
	Descripition *string `json:"descripition,omitempty"`
}

func (o HotfixInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotfixInfo struct{}"
	}

	return strings.Join([]string{"HotfixInfo", string(data)}, " ")
}

type HotfixInfoCommonPatch struct {
	value string
}

type HotfixInfoCommonPatchEnum struct {
	COMMON  HotfixInfoCommonPatch
	CERTAIN HotfixInfoCommonPatch
}

func GetHotfixInfoCommonPatchEnum() HotfixInfoCommonPatchEnum {
	return HotfixInfoCommonPatchEnum{
		COMMON: HotfixInfoCommonPatch{
			value: "common",
		},
		CERTAIN: HotfixInfoCommonPatch{
			value: "certain",
		},
	}
}

func (c HotfixInfoCommonPatch) Value() string {
	return c.value
}

func (c HotfixInfoCommonPatch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HotfixInfoCommonPatch) UnmarshalJSON(b []byte) error {
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
