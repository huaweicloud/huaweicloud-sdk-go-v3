package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeletePortResponse Response Object
type NeutronDeletePortResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeletePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeletePortResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeletePortResponse", string(data)}, " ")
}
