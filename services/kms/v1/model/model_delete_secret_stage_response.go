package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
