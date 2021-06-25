package model

import (
	"encoding/json"

	"strings"
)

type BucketAuthorizedReq struct {
	// 桶名称 ID

	Bucket string `json:"bucket"`
	// 操作标记，0表示取消授权，1表示授权

	Operation string `json:"operation"`
}

func (o BucketAuthorizedReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BucketAuthorizedReq struct{}"
	}

	return strings.Join([]string{"BucketAuthorizedReq", string(data)}, " ")
}
