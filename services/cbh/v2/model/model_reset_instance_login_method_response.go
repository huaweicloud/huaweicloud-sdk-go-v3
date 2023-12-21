package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetInstanceLoginMethodResponse Response Object
type ResetInstanceLoginMethodResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetInstanceLoginMethodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetInstanceLoginMethodResponse struct{}"
	}

	return strings.Join([]string{"ResetInstanceLoginMethodResponse", string(data)}, " ")
}
