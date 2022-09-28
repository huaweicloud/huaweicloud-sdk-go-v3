package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDesktopDetailResponse struct {
	Desktop        *DesktopDetailInfo `json:"desktop,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowDesktopDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDesktopDetailResponse", string(data)}, " ")
}
