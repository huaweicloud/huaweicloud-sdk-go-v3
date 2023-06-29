package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePoolResponse Response Object
type UpdatePoolResponse struct {
	Pool           *PoolResp `json:"pool,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolResponse struct{}"
	}

	return strings.Join([]string{"UpdatePoolResponse", string(data)}, " ")
}
