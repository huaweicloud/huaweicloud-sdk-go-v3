package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListKeypairsResponse struct {
	Body           *[]SimpleKeypair `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListKeypairsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeypairsResponse struct{}"
	}

	return strings.Join([]string{"ListKeypairsResponse", string(data)}, " ")
}
