package model

import (
	"encoding/json"

	"strings"
)

type PlayInfo struct {
	// 播放协议类型<br/>

	PlayType *string `json:"play_type,omitempty"`
	// 播放url<br/>

	Url *string `json:"url,omitempty"`
	// 标记流是否已被加密，取值[0,1] 0表示未加密，1表示已被加密。<br/>

	Encrypted *int32 `json:"encrypted,omitempty"`

	MetaData *MetaData `json:"meta_data,omitempty"`
}

func (o PlayInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PlayInfo struct{}"
	}

	return strings.Join([]string{"PlayInfo", string(data)}, " ")
}
