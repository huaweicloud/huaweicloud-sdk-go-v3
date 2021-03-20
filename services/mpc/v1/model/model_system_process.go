package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SystemProcess struct {
	// 追加转码类型。  取值如下： - SUBTITLE：追加字幕 - AUDIO：追加音频 - VIDEO：追加视频（原有追加新分辨率视频）

	AppendType *SystemProcessAppendType `json:"append_type,omitempty"`

	HlsIndex *ObsObjInfo `json:"hls_index,omitempty"`

	DashIndex *ObsObjInfo `json:"dash_index,omitempty"`
}

func (o SystemProcess) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SystemProcess struct{}"
	}

	return strings.Join([]string{"SystemProcess", string(data)}, " ")
}

type SystemProcessAppendType struct {
	value string
}

type SystemProcessAppendTypeEnum struct {
	SUBTITLE SystemProcessAppendType
	AUDIO    SystemProcessAppendType
	VIDEO    SystemProcessAppendType
}

func GetSystemProcessAppendTypeEnum() SystemProcessAppendTypeEnum {
	return SystemProcessAppendTypeEnum{
		SUBTITLE: SystemProcessAppendType{
			value: "SUBTITLE",
		},
		AUDIO: SystemProcessAppendType{
			value: "AUDIO",
		},
		VIDEO: SystemProcessAppendType{
			value: "VIDEO",
		},
	}
}

func (c SystemProcessAppendType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SystemProcessAppendType) UnmarshalJSON(b []byte) error {
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
