package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppsResponse Response Object
type ListAppsResponse struct {

	// app的总数
	Count *int32 `json:"count,omitempty"`

	// app的列表
	Apps *[]App `json:"apps,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsResponse struct{}"
	}

	return strings.Join([]string{"ListAppsResponse", string(data)}, " ")
}
