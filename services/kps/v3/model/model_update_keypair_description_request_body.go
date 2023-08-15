package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeypairDescriptionRequestBody 更新SSH密钥对描述
type UpdateKeypairDescriptionRequestBody struct {
	Keypair *UpdateKeypairDescriptionReq `json:"keypair"`
}

func (o UpdateKeypairDescriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairDescriptionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeypairDescriptionRequestBody", string(data)}, " ")
}
