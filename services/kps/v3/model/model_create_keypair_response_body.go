package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKeypairResponseBody SSH密钥对的详情
type CreateKeypairResponseBody struct {
	Keypair *CreateKeypairResp `json:"keypair,omitempty"`
}

func (o CreateKeypairResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeypairResponseBody struct{}"
	}

	return strings.Join([]string{"CreateKeypairResponseBody", string(data)}, " ")
}
