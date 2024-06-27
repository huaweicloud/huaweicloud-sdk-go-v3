package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClickHouseInstanceResponse Response Object
type CreateClickHouseInstanceResponse struct {
	Instance *CreateChInstanceInfo `json:"instance,omitempty"`

	// 工作ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClickHouseInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClickHouseInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateClickHouseInstanceResponse", string(data)}, " ")
}
