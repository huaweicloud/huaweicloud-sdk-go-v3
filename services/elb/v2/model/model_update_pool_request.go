package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePoolRequest struct {
	// 后端云服务器组id

	PoolId string `json:"pool_id"`

	Body *UpdatePoolRequestBody `json:"body,omitempty"`
}

func (o UpdatePoolRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePoolRequest struct{}"
	}

	return strings.Join([]string{"UpdatePoolRequest", string(data)}, " ")
}
