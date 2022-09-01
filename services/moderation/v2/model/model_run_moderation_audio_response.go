package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunModerationAudioResponse struct {
	Result         *RunModerationAudioResponseBodyResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int                                   `json:"-"`
}

func (o RunModerationAudioResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModerationAudioResponse struct{}"
	}

	return strings.Join([]string{"RunModerationAudioResponse", string(data)}, " ")
}
