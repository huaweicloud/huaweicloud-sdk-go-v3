package model

import (
	"encoding/json"

	"strings"
)

type SetRecordSetsStatusReq struct {
	// 解析记录状态。  取值范围：  ENABLE：启用解析 DISABLE：暂停解析
	Status string `json:"status"`
}

func (o SetRecordSetsStatusReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetRecordSetsStatusReq struct{}"
	}

	return strings.Join([]string{"SetRecordSetsStatusReq", string(data)}, " ")
}
