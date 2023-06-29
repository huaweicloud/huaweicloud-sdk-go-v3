package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugDataconnectionResponse Response Object
type DebugDataconnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DebugDataconnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugDataconnectionResponse struct{}"
	}

	return strings.Join([]string{"DebugDataconnectionResponse", string(data)}, " ")
}
