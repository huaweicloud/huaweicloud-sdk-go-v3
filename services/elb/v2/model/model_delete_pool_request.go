package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePoolRequest struct {
	// 后端云服务器组id

	PoolId string `json:"pool_id"`
}

func (o DeletePoolRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePoolRequest struct{}"
	}

	return strings.Join([]string{"DeletePoolRequest", string(data)}, " ")
}
