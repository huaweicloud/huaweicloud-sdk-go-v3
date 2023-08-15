package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportEventsResponse Response Object
type ImportEventsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ImportEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportEventsResponse struct{}"
	}

	return strings.Join([]string{"ImportEventsResponse", string(data)}, " ")
}
