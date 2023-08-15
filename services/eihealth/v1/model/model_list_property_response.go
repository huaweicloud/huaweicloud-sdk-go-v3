package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPropertyResponse Response Object
type ListPropertyResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPropertyResponse struct{}"
	}

	return strings.Join([]string{"ListPropertyResponse", string(data)}, " ")
}
