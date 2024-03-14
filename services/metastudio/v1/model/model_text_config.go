package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TextConfig 台词脚本。 > * 最长2000个字符，不含SSML标签字符数。
type TextConfig struct {

	// 台词脚本。  支持两种模式，纯文本模式和标签模式。  ### 纯文本模式 纯文本模式，使用方法，如“大家好，我是人工智大家，是个虚拟主播”。  ### 标签模式 SSML标签的详细定义请参考[文本驱动SSML定义](metastudio_02_0038.xml)。
	Text string `json:"text"`
}

func (o TextConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextConfig struct{}"
	}

	return strings.Join([]string{"TextConfig", string(data)}, " ")
}
