package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCountiesResponse Response Object
type ListCountiesResponse struct {

	// 查询个数，成功的时候返回。
	Count *int32 `json:"count,omitempty"`

	// 区县信息列表，成功的时候返回，具体参见表2。
	Counties       *[]County `json:"counties,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCountiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCountiesResponse struct{}"
	}

	return strings.Join([]string{"ListCountiesResponse", string(data)}, " ")
}
