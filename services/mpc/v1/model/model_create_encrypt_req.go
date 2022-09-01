package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEncryptReq struct {
	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	Encryption *Encryption `json:"encryption,omitempty" xml:"encryption"`

	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`
}

func (o CreateEncryptReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEncryptReq struct{}"
	}

	return strings.Join([]string{"CreateEncryptReq", string(data)}, " ")
}
