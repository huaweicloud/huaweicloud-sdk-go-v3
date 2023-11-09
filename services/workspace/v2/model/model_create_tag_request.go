package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagRequest Request Object
type CreateTagRequest struct {

	// 桌面id。
	DesktopId string `json:"desktop_id"`

	Body *CreateTagReq `json:"body,omitempty"`
}

func (o CreateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTagRequest", string(data)}, " ")
}
