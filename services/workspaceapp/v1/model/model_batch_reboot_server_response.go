package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebootServerResponse Response Object
type BatchRebootServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchRebootServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootServerResponse struct{}"
	}

	return strings.Join([]string{"BatchRebootServerResponse", string(data)}, " ")
}
