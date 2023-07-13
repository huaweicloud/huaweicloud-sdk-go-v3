package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentsResponse Response Object
type ListEnvironmentsResponse struct {

	// 环境总数。
	Count *int32 `json:"count,omitempty"`

	// 环境列表。
	Environments   *[]Environment `json:"environments,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEnvironmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsResponse", string(data)}, " ")
}
