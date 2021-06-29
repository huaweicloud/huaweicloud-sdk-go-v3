package model

import (
	"encoding/json"

	"strings"
)

// 封面信息。
type CoverInfo struct {
	// 封面文件的下载地址。

	CoverUrl *string `json:"cover_url,omitempty"`
}

func (o CoverInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CoverInfo struct{}"
	}

	return strings.Join([]string{"CoverInfo", string(data)}, " ")
}
