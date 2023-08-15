package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretStageResponse Response Object
type ShowSecretStageResponse struct {
	Stage          *Stage `json:"stage,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSecretStageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretStageResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretStageResponse", string(data)}, " ")
}
