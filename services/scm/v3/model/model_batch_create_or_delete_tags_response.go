package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateOrDeleteTagsResponse Response Object
type BatchCreateOrDeleteTagsResponse struct {

	// 请求结果。
	Desc           *string `json:"desc,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateOrDeleteTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteTagsResponse", string(data)}, " ")
}
