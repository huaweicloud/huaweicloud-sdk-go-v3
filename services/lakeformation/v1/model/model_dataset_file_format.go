package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DatasetFileFormat 文件类型
type DatasetFileFormat struct {

	// 文件格式： ROW-行存储文件, TEXT-无格式文本文件, IMAGE-图片文件, AUDIO-音频文件, VIDEO-视频文件, CUSTOM-其他文件
	Format DatasetFileFormatFormat `json:"format"`

	// 文件内容类型。行存文件格式,format为ROW时设置,可选值: CSV-Comma Separated Values文件,JSONL-Json对象行文件,AVRO-AVRO行存文件 图片文件格式，format为IMAGE时设置,可选值: JPG-JPG图片,PNG-PNG图片,TIFF-TIFF图片 音频文件格式，format为AUDIO时设置,可选值: WAV-WAV音频,MP3-MP3音频,FLAC-FLAC音频 视频文件格式，format为VIDEO时设置,可选值: MP4-MP4视频,MOV-MOV视频,AVI-AVI视频
	ContentTypes []string `json:"content_types"`
}

func (o DatasetFileFormat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasetFileFormat struct{}"
	}

	return strings.Join([]string{"DatasetFileFormat", string(data)}, " ")
}

type DatasetFileFormatFormat struct {
	value string
}

type DatasetFileFormatFormatEnum struct {
	ROW    DatasetFileFormatFormat
	TEXT   DatasetFileFormatFormat
	IMAGE  DatasetFileFormatFormat
	AUDIO  DatasetFileFormatFormat
	VIDEO  DatasetFileFormatFormat
	CUSTOM DatasetFileFormatFormat
}

func GetDatasetFileFormatFormatEnum() DatasetFileFormatFormatEnum {
	return DatasetFileFormatFormatEnum{
		ROW: DatasetFileFormatFormat{
			value: "ROW",
		},
		TEXT: DatasetFileFormatFormat{
			value: "TEXT",
		},
		IMAGE: DatasetFileFormatFormat{
			value: "IMAGE",
		},
		AUDIO: DatasetFileFormatFormat{
			value: "AUDIO",
		},
		VIDEO: DatasetFileFormatFormat{
			value: "VIDEO",
		},
		CUSTOM: DatasetFileFormatFormat{
			value: "CUSTOM",
		},
	}
}

func (c DatasetFileFormatFormat) Value() string {
	return c.value
}

func (c DatasetFileFormatFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasetFileFormatFormat) UnmarshalJSON(b []byte) error {
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
