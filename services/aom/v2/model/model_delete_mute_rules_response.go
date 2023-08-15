package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMuteRulesResponse Response Object
type DeleteMuteRulesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMuteRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMuteRulesResponse struct{}"
	}

	return strings.Join([]string{"DeleteMuteRulesResponse", string(data)}, " ")
}
