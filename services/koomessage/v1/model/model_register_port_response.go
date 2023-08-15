package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterPortResponse Response Object
type RegisterPortResponse struct {
	Data           *RegisterPortResponseModel `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o RegisterPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterPortResponse struct{}"
	}

	return strings.Join([]string{"RegisterPortResponse", string(data)}, " ")
}
