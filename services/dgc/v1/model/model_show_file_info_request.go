package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowFileInfoRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *FilePath `json:"body,omitempty"`
}

func (o ShowFileInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowFileInfoRequest", string(data)}, " ")
}
