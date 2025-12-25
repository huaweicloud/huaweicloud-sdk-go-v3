package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipesResponse Response Object
type ListPipesResponse struct {

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 管道列表
	Records        *[]Pipe `json:"records,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPipesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipesResponse struct{}"
	}

	return strings.Join([]string{"ListPipesResponse", string(data)}, " ")
}
