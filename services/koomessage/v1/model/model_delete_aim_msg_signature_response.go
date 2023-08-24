package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAimMsgSignatureResponse Response Object
type DeleteAimMsgSignatureResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAimMsgSignatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAimMsgSignatureResponse struct{}"
	}

	return strings.Join([]string{"DeleteAimMsgSignatureResponse", string(data)}, " ")
}
