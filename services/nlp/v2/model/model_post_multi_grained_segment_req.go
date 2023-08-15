package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PostMultiGrainedSegmentReq This is a auto create Body Object
type PostMultiGrainedSegmentReq struct {

	// 待分词文本，长度为1~64，文本编码为UTF-8。
	Text string `json:"text"`

	// 支持的文本语言类型，目前只支持中文，默认为zh。
	Lang *PostMultiGrainedSegmentReqLang `json:"lang,omitempty"`

	// 分词粒度，1为最细粒度，2为最粗粒度，其它情况默认返回全部粒度分词树结果。
	Granularity *PostMultiGrainedSegmentReqGranularity `json:"granularity,omitempty"`
}

func (o PostMultiGrainedSegmentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostMultiGrainedSegmentReq struct{}"
	}

	return strings.Join([]string{"PostMultiGrainedSegmentReq", string(data)}, " ")
}

type PostMultiGrainedSegmentReqLang struct {
	value string
}

type PostMultiGrainedSegmentReqLangEnum struct {
	ZH PostMultiGrainedSegmentReqLang
}

func GetPostMultiGrainedSegmentReqLangEnum() PostMultiGrainedSegmentReqLangEnum {
	return PostMultiGrainedSegmentReqLangEnum{
		ZH: PostMultiGrainedSegmentReqLang{
			value: "zh",
		},
	}
}

func (c PostMultiGrainedSegmentReqLang) Value() string {
	return c.value
}

func (c PostMultiGrainedSegmentReqLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostMultiGrainedSegmentReqLang) UnmarshalJSON(b []byte) error {
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

type PostMultiGrainedSegmentReqGranularity struct {
	value int32
}

type PostMultiGrainedSegmentReqGranularityEnum struct {
	E_1 PostMultiGrainedSegmentReqGranularity
	E_2 PostMultiGrainedSegmentReqGranularity
	E_0 PostMultiGrainedSegmentReqGranularity
}

func GetPostMultiGrainedSegmentReqGranularityEnum() PostMultiGrainedSegmentReqGranularityEnum {
	return PostMultiGrainedSegmentReqGranularityEnum{
		E_1: PostMultiGrainedSegmentReqGranularity{
			value: 1,
		}, E_2: PostMultiGrainedSegmentReqGranularity{
			value: 2,
		}, E_0: PostMultiGrainedSegmentReqGranularity{
			value: 0,
		},
	}
}

func (c PostMultiGrainedSegmentReqGranularity) Value() int32 {
	return c.value
}

func (c PostMultiGrainedSegmentReqGranularity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostMultiGrainedSegmentReqGranularity) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
