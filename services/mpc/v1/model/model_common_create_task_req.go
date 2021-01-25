package model

import (
	"encoding/json"

	"strings"
)

type CommonCreateTaskReq struct {
	Input  *ObsObjInfo `json:"input,omitempty"`
	Output *ObsObjInfo `json:"output,omitempty"`
	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty"`
	// 是否同步处理, - 0：排队处理 - 1：同步处理  默认值：0
	Sync *int32 `json:"sync,omitempty"`
}

func (o CommonCreateTaskReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CommonCreateTaskReq struct{}"
	}

	return strings.Join([]string{"CommonCreateTaskReq", string(data)}, " ")
}
