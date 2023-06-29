package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersDedicatedHostResponse Response Object
type ListServersDedicatedHostResponse struct {

	// server字段数据结构说明
	Servers        *[]RespServer `json:"servers,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListServersDedicatedHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersDedicatedHostResponse struct{}"
	}

	return strings.Join([]string{"ListServersDedicatedHostResponse", string(data)}, " ")
}
