package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppsBindedToApiV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次返回的API列表

	Auths          *[]ApiAuthInfo `json:"auths,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAppsBindedToApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsBindedToApiV2Response struct{}"
	}

	return strings.Join([]string{"ListAppsBindedToApiV2Response", string(data)}, " ")
}
