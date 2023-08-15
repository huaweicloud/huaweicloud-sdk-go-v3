package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdatePortResponse Response Object
type NeutronUpdatePortResponse struct {
	Port           *NeutronPort `json:"port,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o NeutronUpdatePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdatePortResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdatePortResponse", string(data)}, " ")
}
