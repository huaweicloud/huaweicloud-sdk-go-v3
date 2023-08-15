package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreatePortResponse Response Object
type NeutronCreatePortResponse struct {
	Port           *NeutronPort `json:"port,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o NeutronCreatePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreatePortResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreatePortResponse", string(data)}, " ")
}
