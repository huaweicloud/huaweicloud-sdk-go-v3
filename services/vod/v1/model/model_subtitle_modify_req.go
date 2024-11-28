package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubtitleModifyReq struct {

	// 媒资ID
	AssetId string `json:"asset_id"`

	// 字幕默认语言(字幕必须存在)
	DefaultLanguage *string `json:"default_language,omitempty"`

	// 外挂模式，不传默认取值为0  取值如下： -0：表示添加的字幕会外挂上历史产物 -1：表示添加的字幕不会外挂上历史产物
	RepackageMode *SubtitleModifyReqRepackageMode `json:"repackage_mode,omitempty"`

	// 删除模式，不传默认取值为0  取值如下： -0：表示删除字幕会清除历史产物携带的字幕信息 -1：表示删除字幕不清除历史产物携带的字幕信息
	DeleteMode *SubtitleModifyReqDeleteMode `json:"delete_mode,omitempty"`

	// 需新增或修改的字幕
	AddSubtitles *[]AddSubtitle `json:"add_subtitles,omitempty"`

	// 需删除的字幕，language不能与add_subtitles重复
	DeleteSubtitles *[]DeleteSubtitle `json:"delete_subtitles,omitempty"`
}

func (o SubtitleModifyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubtitleModifyReq struct{}"
	}

	return strings.Join([]string{"SubtitleModifyReq", string(data)}, " ")
}

type SubtitleModifyReqRepackageMode struct {
	value string
}

type SubtitleModifyReqRepackageModeEnum struct {
	E_0 SubtitleModifyReqRepackageMode
	E_1 SubtitleModifyReqRepackageMode
}

func GetSubtitleModifyReqRepackageModeEnum() SubtitleModifyReqRepackageModeEnum {
	return SubtitleModifyReqRepackageModeEnum{
		E_0: SubtitleModifyReqRepackageMode{
			value: "0",
		},
		E_1: SubtitleModifyReqRepackageMode{
			value: "1",
		},
	}
}

func (c SubtitleModifyReqRepackageMode) Value() string {
	return c.value
}

func (c SubtitleModifyReqRepackageMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubtitleModifyReqRepackageMode) UnmarshalJSON(b []byte) error {
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

type SubtitleModifyReqDeleteMode struct {
	value string
}

type SubtitleModifyReqDeleteModeEnum struct {
	E_0 SubtitleModifyReqDeleteMode
	E_1 SubtitleModifyReqDeleteMode
}

func GetSubtitleModifyReqDeleteModeEnum() SubtitleModifyReqDeleteModeEnum {
	return SubtitleModifyReqDeleteModeEnum{
		E_0: SubtitleModifyReqDeleteMode{
			value: "0",
		},
		E_1: SubtitleModifyReqDeleteMode{
			value: "1",
		},
	}
}

func (c SubtitleModifyReqDeleteMode) Value() string {
	return c.value
}

func (c SubtitleModifyReqDeleteMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubtitleModifyReqDeleteMode) UnmarshalJSON(b []byte) error {
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
