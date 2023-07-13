package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsedDesktopInfoRequest Request Object
type ListUsedDesktopInfoRequest struct {
	Body *ListUsedDesktopInfoReq `json:"body,omitempty"`
}

func (o ListUsedDesktopInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsedDesktopInfoRequest struct{}"
	}

	return strings.Join([]string{"ListUsedDesktopInfoRequest", string(data)}, " ")
}
