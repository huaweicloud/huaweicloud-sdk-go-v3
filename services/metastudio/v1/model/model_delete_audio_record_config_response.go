package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAudioRecordConfigResponse Response Object
type DeleteAudioRecordConfigResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAudioRecordConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAudioRecordConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteAudioRecordConfigResponse", string(data)}, " ")
}
