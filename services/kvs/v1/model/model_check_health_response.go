package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckHealthResponse Response Object
type CheckHealthResponse struct {
	HttpStatusCode int `bson:"-"`
}

func (o CheckHealthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckHealthResponse struct{}"
	}

	return strings.Join([]string{"CheckHealthResponse", string(data)}, " ")
}
