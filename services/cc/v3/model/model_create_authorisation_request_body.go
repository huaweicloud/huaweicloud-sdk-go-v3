package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorisationRequestBody 创建授权的详细信息。
type CreateAuthorisationRequestBody struct {
	Authorisation *CreateAuthorisation `json:"authorisation"`
}

func (o CreateAuthorisationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorisationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAuthorisationRequestBody", string(data)}, " ")
}
