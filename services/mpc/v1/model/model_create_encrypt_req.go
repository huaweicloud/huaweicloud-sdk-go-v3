package model

import (
	"encoding/json"

	"strings"
)

type CreateEncryptReq struct {
	Input      *ObsObjInfo `json:"input,omitempty"`
	Output     *ObsObjInfo `json:"output,omitempty"`
	Encryption *Encryption `json:"encryption,omitempty"`
	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty"`
}

func (o CreateEncryptReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEncryptReq struct{}"
	}

	return strings.Join([]string{"CreateEncryptReq", string(data)}, " ")
}
