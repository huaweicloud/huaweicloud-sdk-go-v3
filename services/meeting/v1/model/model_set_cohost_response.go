package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetCohostResponse Response Object
type SetCohostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetCohostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCohostResponse struct{}"
	}

	return strings.Join([]string{"SetCohostResponse", string(data)}, " ")
}
