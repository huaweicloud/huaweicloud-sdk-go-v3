package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Slot struct {

	// 实体文本。
	Word string `json:"word"`

	// 实体类型。对于每个意图类别所支持的实体类型分别为： weather：date(日期)，time(时间)，location(位置) time：location(位置)，timezone(时区) news：genre(风格) joke：genre(风格) translation：content(内容) notification：content(内容)，date(日期)，time(时间)，singer(歌手) alarm：date(日期)，time:(时间) music：singer(歌手)，song(歌曲)，content(内容)
	Tag string `json:"tag"`

	// 实体文本在待分析文本中的起始位置。
	Offset int32 `json:"offset"`

	// 实体文本长度。
	Length int32 `json:"length"`

	// 同义词或者其他标准表达的词，默认为原始的word。
	NormalizedWord string `json:"normalized_word"`
}

func (o Slot) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Slot struct{}"
	}

	return strings.Join([]string{"Slot", string(data)}, " ")
}
