package model

import (
	"encoding/json"

	"strings"
)

type DeleteFacesBatchReq struct {
	// 过滤条件，参考[filter语法](https://support.huaweicloud.com/api-face/face_02_0014.html)。

	Filter string `json:"filter"`
}

func (o DeleteFacesBatchReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFacesBatchReq struct{}"
	}

	return strings.Join([]string{"DeleteFacesBatchReq", string(data)}, " ")
}
