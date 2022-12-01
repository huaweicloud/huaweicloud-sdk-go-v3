package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShrinkClusterResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkClusterResponse struct{}"
	}

	return strings.Join([]string{"ShrinkClusterResponse", string(data)}, " ")
}
