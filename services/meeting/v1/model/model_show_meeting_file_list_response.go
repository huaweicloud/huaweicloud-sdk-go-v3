package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMeetingFileListResponse Response Object
type ShowMeetingFileListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowMeetingFileListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeetingFileListResponse struct{}"
	}

	return strings.Join([]string{"ShowMeetingFileListResponse", string(data)}, " ")
}
