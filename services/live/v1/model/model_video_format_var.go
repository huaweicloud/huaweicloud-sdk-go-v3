package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoFormatVar struct {
	value string
}

type VideoFormatVarEnum struct {
	FLV VideoFormatVar
	HLS VideoFormatVar
	MP4 VideoFormatVar
}

func GetVideoFormatVarEnum() VideoFormatVarEnum {
	return VideoFormatVarEnum{
		FLV: VideoFormatVar{
			value: "FLV",
		},
		HLS: VideoFormatVar{
			value: "HLS",
		},
		MP4: VideoFormatVar{
			value: "MP4",
		},
	}
}

func (c VideoFormatVar) Value() string {
	return c.value
}

func (c VideoFormatVar) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoFormatVar) UnmarshalJSON(b []byte) error {
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
