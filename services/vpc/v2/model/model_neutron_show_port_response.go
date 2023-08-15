package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowPortResponse Response Object
type NeutronShowPortResponse struct {
	Port           *NeutronPort `json:"port,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o NeutronShowPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowPortResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowPortResponse", string(data)}, " ")
}
