package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnvResponse Response Object
type UpdateEnvResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnvResponse", string(data)}, " ")
}
