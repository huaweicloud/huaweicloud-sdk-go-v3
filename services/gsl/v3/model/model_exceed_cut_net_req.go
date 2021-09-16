package model

import (
	"encoding/json"

	"strings"
)

type ExceedCutNetReq struct {
	// 操作类型(1:设置达量断网域值，2：取消达量断网域值)

	Action int32 `json:"action"`
	// 阈值,只能是0,-1,正整数，-1表示无限制，0表示有上网流量产生就会立即断网，取消达量断网功能时可不传，单位MB

	Quota *string `json:"quota,omitempty"`
}

func (o ExceedCutNetReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExceedCutNetReq struct{}"
	}

	return strings.Join([]string{"ExceedCutNetReq", string(data)}, " ")
}
