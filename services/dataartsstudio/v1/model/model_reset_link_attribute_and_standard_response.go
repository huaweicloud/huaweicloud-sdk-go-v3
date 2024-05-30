package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetLinkAttributeAndStandardResponse Response Object
type ResetLinkAttributeAndStandardResponse struct {
	Data           *ResetLinkAttributeAndStandardResultData `json:"data,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ResetLinkAttributeAndStandardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetLinkAttributeAndStandardResponse struct{}"
	}

	return strings.Join([]string{"ResetLinkAttributeAndStandardResponse", string(data)}, " ")
}
