package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseConferenceResponse Response Object
type PauseConferenceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PauseConferenceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseConferenceResponse struct{}"
	}

	return strings.Join([]string{"PauseConferenceResponse", string(data)}, " ")
}
