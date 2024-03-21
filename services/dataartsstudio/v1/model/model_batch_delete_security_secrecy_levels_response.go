package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecuritySecrecyLevelsResponse Response Object
type BatchDeleteSecuritySecrecyLevelsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSecuritySecrecyLevelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecuritySecrecyLevelsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecuritySecrecyLevelsResponse", string(data)}, " ")
}
