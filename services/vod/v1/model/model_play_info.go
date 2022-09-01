package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PlayInfo struct {

	// 播放协议类型。  取值如下： - hls - dash - mp4
	PlayType *string `json:"play_type,omitempty" xml:"play_type"`

	// 播放URL。
	Url *string `json:"url,omitempty" xml:"url"`

	// 标记流是否已被加密。  取值如下： - 0：表示未加密。 - 1：表示已被加密。  默认值：0。
	Encrypted *int32 `json:"encrypted,omitempty" xml:"encrypted"`

	MetaData *MetaData `json:"meta_data,omitempty" xml:"meta_data"`
}

func (o PlayInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlayInfo struct{}"
	}

	return strings.Join([]string{"PlayInfo", string(data)}, " ")
}
