package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShadowsResponse Response Object
type ListShadowsResponse struct {
	Body           *[]ShadowService `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListShadowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShadowsResponse struct{}"
	}

	return strings.Join([]string{"ListShadowsResponse", string(data)}, " ")
}
