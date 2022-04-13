package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
