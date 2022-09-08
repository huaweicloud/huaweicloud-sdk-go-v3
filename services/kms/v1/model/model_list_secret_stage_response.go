package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecretStageResponse struct {
	Stage          *Stage `json:"stage,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecretStageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretStageResponse struct{}"
	}

	return strings.Join([]string{"ListSecretStageResponse", string(data)}, " ")
}
