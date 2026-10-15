package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChapterConfigReq struct {

	// 是否启用该章节。BASIC_INFO章节始终为true。
	Enabled bool `json:"enabled"`

	// 子章节配置，key为子章节编码，value为是否启用。无子章节的章节可不传此字段。
	SubChapters map[string]bool `json:"sub_chapters,omitempty"`
}

func (o ChapterConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChapterConfigReq struct{}"
	}

	return strings.Join([]string{"ChapterConfigReq", string(data)}, " ")
}
