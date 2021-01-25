package model

import (
	"encoding/json"

	"strings"
)

type Encryption struct {
	HlsEncrypt *HlsEncrypt `json:"hls_encrypt,omitempty"`
	Multidrm   *Multidrm   `json:"multidrm,omitempty"`
	// 加密预览时长, 单位秒(S), 0 - preview_duration之间的内容不加密
	PreviewDuration *int32 `json:"preview_duration,omitempty"`
}

func (o Encryption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Encryption struct{}"
	}

	return strings.Join([]string{"Encryption", string(data)}, " ")
}
