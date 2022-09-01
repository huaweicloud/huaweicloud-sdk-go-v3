package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 分词post请求体
type SegmentRequest struct {

	// 待分词文本，长度为1~512，文本编码为UTF-8。
	Text string `json:"text" xml:"text"`

	// 是否开启词性标注功能，1为开启，0为关闭，默认为关闭。
	PosSwitch *SegmentRequestPosSwitch `json:"pos_switch,omitempty" xml:"pos_switch"`

	// 支持的文本语言类型，目前支持中文（zh）和英文（en），默认为中文。
	Lang *SegmentRequestLang `json:"lang,omitempty" xml:"lang"`

	// 支持的分词规范。 中文分词标准目前支持PKU（北大分词标准）、CTB（宾州中文树库标准），默认为PKU。 英文分词标准默认为Penn TreeBank（宾州树库标准），不需要传入该参数。
	Criterion *SegmentRequestCriterion `json:"criterion,omitempty" xml:"criterion"`
}

func (o SegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SegmentRequest struct{}"
	}

	return strings.Join([]string{"SegmentRequest", string(data)}, " ")
}

type SegmentRequestPosSwitch struct {
	value int32
}

type SegmentRequestPosSwitchEnum struct {
	E_0 SegmentRequestPosSwitch
	E_1 SegmentRequestPosSwitch
}

func GetSegmentRequestPosSwitchEnum() SegmentRequestPosSwitchEnum {
	return SegmentRequestPosSwitchEnum{
		E_0: SegmentRequestPosSwitch{
			value: 0,
		}, E_1: SegmentRequestPosSwitch{
			value: 1,
		},
	}
}

func (c SegmentRequestPosSwitch) Value() int32 {
	return c.value
}

func (c SegmentRequestPosSwitch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SegmentRequestPosSwitch) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type SegmentRequestLang struct {
	value string
}

type SegmentRequestLangEnum struct {
	ZH SegmentRequestLang
	EN SegmentRequestLang
}

func GetSegmentRequestLangEnum() SegmentRequestLangEnum {
	return SegmentRequestLangEnum{
		ZH: SegmentRequestLang{
			value: "zh",
		},
		EN: SegmentRequestLang{
			value: "en",
		},
	}
}

func (c SegmentRequestLang) Value() string {
	return c.value
}

func (c SegmentRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SegmentRequestLang) UnmarshalJSON(b []byte) error {
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

type SegmentRequestCriterion struct {
	value string
}

type SegmentRequestCriterionEnum struct {
	PKU SegmentRequestCriterion
	CTB SegmentRequestCriterion
}

func GetSegmentRequestCriterionEnum() SegmentRequestCriterionEnum {
	return SegmentRequestCriterionEnum{
		PKU: SegmentRequestCriterion{
			value: "PKU",
		},
		CTB: SegmentRequestCriterion{
			value: "CTB",
		},
	}
}

func (c SegmentRequestCriterion) Value() string {
	return c.value
}

func (c SegmentRequestCriterion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SegmentRequestCriterion) UnmarshalJSON(b []byte) error {
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
