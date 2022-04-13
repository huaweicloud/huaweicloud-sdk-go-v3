package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新SSH密钥对描述
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
