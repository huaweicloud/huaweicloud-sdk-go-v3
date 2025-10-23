package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportedRegionResponse Response Object
type ListSupportedRegionResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSupportedRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedRegionResponse struct{}"
	}

	return strings.Join([]string{"ListSupportedRegionResponse", string(data)}, " ")
}
