package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTerminalsBindingDesktopsResponse Response Object
type CreateTerminalsBindingDesktopsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTerminalsBindingDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTerminalsBindingDesktopsResponse struct{}"
	}

	return strings.Join([]string{"CreateTerminalsBindingDesktopsResponse", string(data)}, " ")
}
