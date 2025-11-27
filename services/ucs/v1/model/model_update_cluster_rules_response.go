package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterRulesResponse Response Object
type UpdateClusterRulesResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateClusterRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterRulesResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterRulesResponse", string(data)}, " ")
}
