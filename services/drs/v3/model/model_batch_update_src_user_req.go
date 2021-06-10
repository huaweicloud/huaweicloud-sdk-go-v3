package model

import (
	"encoding/json"

	"strings"
)

// 批量更新迁移用户请求体
type BatchUpdateSrcUserReq struct {
	// 批量更新迁移用户请求列表

	Jobs []UpdateUserReq `json:"jobs"`
}

func (o BatchUpdateSrcUserReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchUpdateSrcUserReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateSrcUserReq", string(data)}, " ")
}
