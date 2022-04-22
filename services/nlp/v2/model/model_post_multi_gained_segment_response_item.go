package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 多粒度分词结果中的词汇节点
type PostMultiGainedSegmentResponseItem struct {

	// 当前节点对应的文本内容
	Content string `json:"content"`

	// 文本类型，取值如下： WORD-词汇类型 CHAR-字符类型
	Type PostMultiGainedSegmentResponseItemType `json:"type"`

	// 当前节点的子节点列表
	SubContents *[]PostMultiGainedSegmentResponseItemCopy `json:"sub_contents,omitempty"`
}

func (o PostMultiGainedSegmentResponseItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostMultiGainedSegmentResponseItem struct{}"
	}

	return strings.Join([]string{"PostMultiGainedSegmentResponseItem", string(data)}, " ")
}

type PostMultiGainedSegmentResponseItemType struct {
	value string
}

type PostMultiGainedSegmentResponseItemTypeEnum struct {
	WORD PostMultiGainedSegmentResponseItemType
	CHAR PostMultiGainedSegmentResponseItemType
}

func GetPostMultiGainedSegmentResponseItemTypeEnum() PostMultiGainedSegmentResponseItemTypeEnum {
	return PostMultiGainedSegmentResponseItemTypeEnum{
		WORD: PostMultiGainedSegmentResponseItemType{
			value: "WORD",
		},
		CHAR: PostMultiGainedSegmentResponseItemType{
			value: "CHAR",
		},
	}
}

func (c PostMultiGainedSegmentResponseItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostMultiGainedSegmentResponseItemType) UnmarshalJSON(b []byte) error {
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
