package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretStageResponse Response Object
type DeleteSecretStageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecretStageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretStageResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretStageResponse", string(data)}, " ")
}
