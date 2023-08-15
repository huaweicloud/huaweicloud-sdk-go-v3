package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPartitionsResponse Response Object
type AddPartitionsResponse struct {
	Body           *[]Partition `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o AddPartitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPartitionsResponse struct{}"
	}

	return strings.Join([]string{"AddPartitionsResponse", string(data)}, " ")
}
