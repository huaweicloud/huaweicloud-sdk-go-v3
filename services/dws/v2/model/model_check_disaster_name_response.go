package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDisasterNameResponse Response Object
type CheckDisasterNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckDisasterNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDisasterNameResponse struct{}"
	}

	return strings.Join([]string{"CheckDisasterNameResponse", string(data)}, " ")
}
