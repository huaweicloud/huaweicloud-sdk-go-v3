package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowChildNumResponse Response Object
type ShowChildNumResponse struct {

	// 子任务的数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowChildNumResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowChildNumResponse struct{}"
	}

	return strings.Join([]string{"ShowChildNumResponse", string(data)}, " ")
}
