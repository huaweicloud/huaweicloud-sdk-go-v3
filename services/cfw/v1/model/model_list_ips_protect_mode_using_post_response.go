package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIpsProtectModeUsingPostResponse struct {

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 防护状态
	Status         *int32 `json:"status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIpsProtectModeUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsProtectModeUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ListIpsProtectModeUsingPostResponse", string(data)}, " ")
}
