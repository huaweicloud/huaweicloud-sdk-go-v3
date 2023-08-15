package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSystemEventResponse Response Object
type CreateSystemEventResponse struct {
	Systemevent    *Event `json:"systemevent,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateSystemEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSystemEventResponse struct{}"
	}

	return strings.Join([]string{"CreateSystemEventResponse", string(data)}, " ")
}
