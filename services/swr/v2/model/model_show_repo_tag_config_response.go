package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepoTagConfigResponse Response Object
type ShowRepoTagConfigResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepoTagConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoTagConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowRepoTagConfigResponse", string(data)}, " ")
}
