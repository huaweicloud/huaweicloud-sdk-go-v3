package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactoryEnvResponse Response Object
type CreateFactoryEnvResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	StatusCode     *int32 `json:"status_code,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateFactoryEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactoryEnvResponse struct{}"
	}

	return strings.Join([]string{"CreateFactoryEnvResponse", string(data)}, " ")
}
