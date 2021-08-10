package model

import (
	"encoding/json"

	"strings"
)

type DeleteFacesBatchReq struct {
	// 过滤条件，参考[filter语法](zh-cn_topic_0130807048.xml)。

	Filter string `json:"filter"`
}

func (o DeleteFacesBatchReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFacesBatchReq struct{}"
	}

	return strings.Join([]string{"DeleteFacesBatchReq", string(data)}, " ")
}
