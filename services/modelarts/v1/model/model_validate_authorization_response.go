package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateAuthorizationResponse Response Object
type ValidateAuthorizationResponse struct {

	// **参数解释**：鉴权结果。
	Results        *[]ValidateAuthResults `json:"results,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ValidateAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"ValidateAuthorizationResponse", string(data)}, " ")
}
