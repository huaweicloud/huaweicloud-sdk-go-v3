package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 多粒度分词结果中的词汇节点
type PostMultiGainedSegmentResponseItemCopy struct {

	// 当前节点对应的文本内容
	Content string `json:"content"`

	// 文本类型，取值如下： WORD-词汇类型 CHAR-字符类型
	Type PostMultiGainedSegmentResponseItemCopyType `json:"type"`

	// 当前节点的子节点列表
	SubContents *[]PostMultiGainedSegmentResponseItem `json:"sub_contents,omitempty"`
}

func (o PostMultiGainedSegmentResponseItemCopy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostMultiGainedSegmentResponseItemCopy struct{}"
	}

	return strings.Join([]string{"PostMultiGainedSegmentResponseItemCopy", string(data)}, " ")
}

type PostMultiGainedSegmentResponseItemCopyType struct {
	value string
}

type PostMultiGainedSegmentResponseItemCopyTypeEnum struct {
	WORD PostMultiGainedSegmentResponseItemCopyType
	CHAR PostMultiGainedSegmentResponseItemCopyType
}

func GetPostMultiGainedSegmentResponseItemCopyTypeEnum() PostMultiGainedSegmentResponseItemCopyTypeEnum {
	return PostMultiGainedSegmentResponseItemCopyTypeEnum{
		WORD: PostMultiGainedSegmentResponseItemCopyType{
			value: "WORD",
		},
		CHAR: PostMultiGainedSegmentResponseItemCopyType{
			value: "CHAR",
		},
	}
}

func (c PostMultiGainedSegmentResponseItemCopyType) Value() string {
	return c.value
}

func (c PostMultiGainedSegmentResponseItemCopyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostMultiGainedSegmentResponseItemCopyType) UnmarshalJSON(b []byte) error {
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
