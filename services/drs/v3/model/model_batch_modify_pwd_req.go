package model

import (
	"encoding/json"

	"strings"
)

//
type BatchModifyPwdReq struct {
	// 批量修改数据库密码信息列表

	Jobs []ModifyPwdEndPoint `json:"jobs"`
}

func (o BatchModifyPwdReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchModifyPwdReq struct{}"
	}

	return strings.Join([]string{"BatchModifyPwdReq", string(data)}, " ")
}
