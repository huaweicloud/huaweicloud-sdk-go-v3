package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckMd5DuplicationRequest struct {
	// 文件大小

	Size int64 `json:"size"`
	// 文件MD5

	Md5 string `json:"md5"`
}

func (o CheckMd5DuplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckMd5DuplicationRequest struct{}"
	}

	return strings.Join([]string{"CheckMd5DuplicationRequest", string(data)}, " ")
}
