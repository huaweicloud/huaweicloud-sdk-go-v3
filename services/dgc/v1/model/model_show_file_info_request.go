package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowFileInfoRequest struct {
	Body *FilePath `json:"body,omitempty"`
}

func (o ShowFileInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowFileInfoRequest", string(data)}, " ")
}
