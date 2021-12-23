package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSecretStageResponse struct {
	Stage          *Stage `json:"stage,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateSecretStageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretStageResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretStageResponse", string(data)}, " ")
}
