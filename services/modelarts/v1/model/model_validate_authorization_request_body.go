package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidateAuthorizationRequestBody struct {

	// **参数解释**：工作空间鉴权请求体。
	Requests []AuthRequests `json:"requests"`
}

func (o ValidateAuthorizationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateAuthorizationRequestBody struct{}"
	}

	return strings.Join([]string{"ValidateAuthorizationRequestBody", string(data)}, " ")
}
