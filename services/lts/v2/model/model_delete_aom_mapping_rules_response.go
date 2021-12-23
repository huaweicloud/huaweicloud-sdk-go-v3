package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAomMappingRulesResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteAomMappingRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAomMappingRulesResponse struct{}"
	}

	return strings.Join([]string{"DeleteAomMappingRulesResponse", string(data)}, " ")
}
