package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoffWebCliResponse Response Object
type LogoffWebCliResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LogoffWebCliResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffWebCliResponse struct{}"
	}

	return strings.Join([]string{"LogoffWebCliResponse", string(data)}, " ")
}
