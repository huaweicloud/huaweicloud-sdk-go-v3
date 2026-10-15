package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChapterConfig struct {

	// 是否启用该章节。
	Enabled bool `json:"enabled"`

	// 子章节配置，key为子章节编码，value为是否启用。
	SubChapters map[string]bool `json:"sub_chapters,omitempty"`
}

func (o ChapterConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChapterConfig struct{}"
	}

	return strings.Join([]string{"ChapterConfig", string(data)}, " ")
}
