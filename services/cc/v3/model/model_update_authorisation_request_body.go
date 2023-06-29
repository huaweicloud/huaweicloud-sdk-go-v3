package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthorisationRequestBody 更新授权的详细信息。
type UpdateAuthorisationRequestBody struct {
	Authorisation *UpdateAuthorisation `json:"authorisation"`
}

func (o UpdateAuthorisationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthorisationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAuthorisationRequestBody", string(data)}, " ")
}
