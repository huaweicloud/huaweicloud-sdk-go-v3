package model

import (
	"encoding/json"

	"strings"
)

type CreateRemuxTaskReq struct {
	Input  *ObsObjInfo `json:"input,omitempty"`
	Output *ObsObjInfo `json:"output,omitempty"`
	// 用户自定义数据。
	UserData    *string           `json:"user_data,omitempty"`
	OutputParam *RemuxOutputParam `json:"output_param,omitempty"`
	// 提供给mpe通知回调用的的url
	NotifyUrl *string `json:"notify_url,omitempty"`
}

func (o CreateRemuxTaskReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRemuxTaskReq struct{}"
	}

	return strings.Join([]string{"CreateRemuxTaskReq", string(data)}, " ")
}
