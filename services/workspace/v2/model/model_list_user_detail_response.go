package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserDetailResponse Response Object
type ListUserDetailResponse struct {
	UserDetail     *UserDetail `json:"user_detail,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListUserDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserDetailResponse struct{}"
	}

	return strings.Join([]string{"ListUserDetailResponse", string(data)}, " ")
}
