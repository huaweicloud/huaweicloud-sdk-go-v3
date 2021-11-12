package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type KeyWordsStat struct {
	// 关键词。

	Keyword string `json:"keyword"`
	// 关键词频次。

	Freq int64 `json:"freq"`
}

func (o KeyWordsStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyWordsStat struct{}"
	}

	return strings.Join([]string{"KeyWordsStat", string(data)}, " ")
}
