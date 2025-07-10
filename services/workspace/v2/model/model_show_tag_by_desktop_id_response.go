package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagByDesktopIdResponse Response Object
type ShowTagByDesktopIdResponse struct {

	// 标签。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTagByDesktopIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagByDesktopIdResponse struct{}"
	}

	return strings.Join([]string{"ShowTagByDesktopIdResponse", string(data)}, " ")
}
