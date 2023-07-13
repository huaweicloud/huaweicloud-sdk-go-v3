package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AudioInputBody 音频数据的输入
type AudioInputBody struct {

	// 音频url地址。
	Url string `json:"url"`
}

func (o AudioInputBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioInputBody struct{}"
	}

	return strings.Join([]string{"AudioInputBody", string(data)}, " ")
}
