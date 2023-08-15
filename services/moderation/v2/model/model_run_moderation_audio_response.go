package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunModerationAudioResponse Response Object
type RunModerationAudioResponse struct {
	Result         *RunModerationAudioResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o RunModerationAudioResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModerationAudioResponse struct{}"
	}

	return strings.Join([]string{"RunModerationAudioResponse", string(data)}, " ")
}
