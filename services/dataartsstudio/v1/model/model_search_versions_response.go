package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchVersionsResponse Response Object
type SearchVersionsResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SearchVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchVersionsResponse struct{}"
	}

	return strings.Join([]string{"SearchVersionsResponse", string(data)}, " ")
}
