package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelationsResponse Response Object
type ListRelationsResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListRelationsResponse", string(data)}, " ")
}
