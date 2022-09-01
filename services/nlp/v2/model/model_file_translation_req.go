package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type FileTranslationReq struct {

	// 存放在OBS的文档文件路径，私密文件推荐使用临时授权URL调用服务，如何获取OBS文件URL和临时授权URL请参见配置OBS访问权限（https://support.huaweicloud.com/api-nlp/nlp_03_0080.html）。OBS的region要和请求服务的region保持一致，region不一致则OBS不可用，即使obs是公开访问权限。
	Url string `json:"url" xml:"url"`

	// 翻译原语言，文档翻译服务当前仅支持中英互译。
	From FileTranslationReqFrom `json:"from" xml:"from"`

	// 翻译目标语言，文档翻译服务当前仅支持中英互译。
	To FileTranslationReqTo `json:"to" xml:"to"`

	// 文档格式，当前仅支持翻译“docx”、“pptx”和“txt”格式的文档。
	Type FileTranslationReqType `json:"type" xml:"type"`
}

func (o FileTranslationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileTranslationReq struct{}"
	}

	return strings.Join([]string{"FileTranslationReq", string(data)}, " ")
}

type FileTranslationReqFrom struct {
	value string
}

type FileTranslationReqFromEnum struct {
	ZH FileTranslationReqFrom
	EN FileTranslationReqFrom
}

func GetFileTranslationReqFromEnum() FileTranslationReqFromEnum {
	return FileTranslationReqFromEnum{
		ZH: FileTranslationReqFrom{
			value: "zh",
		},
		EN: FileTranslationReqFrom{
			value: "en",
		},
	}
}

func (c FileTranslationReqFrom) Value() string {
	return c.value
}

func (c FileTranslationReqFrom) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileTranslationReqFrom) UnmarshalJSON(b []byte) error {
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

type FileTranslationReqTo struct {
	value string
}

type FileTranslationReqToEnum struct {
	ZH FileTranslationReqTo
	EN FileTranslationReqTo
}

func GetFileTranslationReqToEnum() FileTranslationReqToEnum {
	return FileTranslationReqToEnum{
		ZH: FileTranslationReqTo{
			value: "zh",
		},
		EN: FileTranslationReqTo{
			value: "en",
		},
	}
}

func (c FileTranslationReqTo) Value() string {
	return c.value
}

func (c FileTranslationReqTo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileTranslationReqTo) UnmarshalJSON(b []byte) error {
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

type FileTranslationReqType struct {
	value string
}

type FileTranslationReqTypeEnum struct {
	DOCX FileTranslationReqType
	PPTX FileTranslationReqType
	TXT  FileTranslationReqType
}

func GetFileTranslationReqTypeEnum() FileTranslationReqTypeEnum {
	return FileTranslationReqTypeEnum{
		DOCX: FileTranslationReqType{
			value: "docx",
		},
		PPTX: FileTranslationReqType{
			value: "pptx",
		},
		TXT: FileTranslationReqType{
			value: "txt",
		},
	}
}

func (c FileTranslationReqType) Value() string {
	return c.value
}

func (c FileTranslationReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileTranslationReqType) UnmarshalJSON(b []byte) error {
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
