package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HotfixInfoResult 热补丁信息
type HotfixInfoResult struct {

	// 热补丁版本
	Version *string `json:"version,omitempty"`

	// 通用非通用信息,common=通用补丁,certain=定制补丁
	CommonPatch *HotfixInfoResultCommonPatch `json:"common_patch,omitempty"`

	// 是否和备份相关
	BackupSensitive *bool `json:"backup_sensitive,omitempty"`

	// 补丁的描述信息
	Descripition *string `json:"descripition,omitempty"`
}

func (o HotfixInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotfixInfoResult struct{}"
	}

	return strings.Join([]string{"HotfixInfoResult", string(data)}, " ")
}

type HotfixInfoResultCommonPatch struct {
	value string
}

type HotfixInfoResultCommonPatchEnum struct {
	COMMON  HotfixInfoResultCommonPatch
	CERTAIN HotfixInfoResultCommonPatch
}

func GetHotfixInfoResultCommonPatchEnum() HotfixInfoResultCommonPatchEnum {
	return HotfixInfoResultCommonPatchEnum{
		COMMON: HotfixInfoResultCommonPatch{
			value: "common",
		},
		CERTAIN: HotfixInfoResultCommonPatch{
			value: "certain",
		},
	}
}

func (c HotfixInfoResultCommonPatch) Value() string {
	return c.value
}

func (c HotfixInfoResultCommonPatch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HotfixInfoResultCommonPatch) UnmarshalJSON(b []byte) error {
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
