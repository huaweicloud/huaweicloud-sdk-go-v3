package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKeypairResponse Response Object
type CreateKeypairResponse struct {
	Keypair        *CreateKeypairResp `json:"keypair,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeypairResponse struct{}"
	}

	return strings.Join([]string{"CreateKeypairResponse", string(data)}, " ")
}
