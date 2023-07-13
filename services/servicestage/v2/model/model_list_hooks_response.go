package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHooksResponse Response Object
type ListHooksResponse struct {

	// hook列表。
	Hooks          *[]Hook `json:"hooks,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHooksResponse struct{}"
	}

	return strings.Join([]string{"ListHooksResponse", string(data)}, " ")
}
