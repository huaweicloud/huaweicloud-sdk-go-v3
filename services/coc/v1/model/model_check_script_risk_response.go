package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckScriptRiskResponse Response Object
type CheckScriptRiskResponse struct {
	Data           *CheckScriptRiskResData `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CheckScriptRiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckScriptRiskResponse struct{}"
	}

	return strings.Join([]string{"CheckScriptRiskResponse", string(data)}, " ")
}
