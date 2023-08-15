package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogDumpObsResponse Response Object
type CreateLogDumpObsResponse struct {

	// 转储id。
	LogDumpObsId   *string `json:"log_dump_obs_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogDumpObsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogDumpObsResponse struct{}"
	}

	return strings.Join([]string{"CreateLogDumpObsResponse", string(data)}, " ")
}
