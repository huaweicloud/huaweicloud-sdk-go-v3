package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiApiKeyResponse Response Object
type DeleteAiApiKeyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAiApiKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiApiKeyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAiApiKeyResponse", string(data)}, " ")
}
