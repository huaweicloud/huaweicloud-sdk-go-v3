package model

import (
	"encoding/json"

	"strings"
)

type CoverInfo struct {
	// 封面文件的下载地址<br/>

	CoverUrl *string `json:"cover_url,omitempty"`
}

func (o CoverInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CoverInfo struct{}"
	}

	return strings.Join([]string{"CoverInfo", string(data)}, " ")
}
