package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelExtractAudioTaskResponse Response Object
type CancelExtractAudioTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelExtractAudioTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelExtractAudioTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelExtractAudioTaskResponse", string(data)}, " ")
}
