package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopPoolAuthorizedObjectsResponse Response Object
type CreateDesktopPoolAuthorizedObjectsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDesktopPoolAuthorizedObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopPoolAuthorizedObjectsResponse struct{}"
	}

	return strings.Join([]string{"CreateDesktopPoolAuthorizedObjectsResponse", string(data)}, " ")
}
