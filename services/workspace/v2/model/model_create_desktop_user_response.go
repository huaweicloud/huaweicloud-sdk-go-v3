package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopUserResponse Response Object
type CreateDesktopUserResponse struct {

	// 用户id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDesktopUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopUserResponse struct{}"
	}

	return strings.Join([]string{"CreateDesktopUserResponse", string(data)}, " ")
}
