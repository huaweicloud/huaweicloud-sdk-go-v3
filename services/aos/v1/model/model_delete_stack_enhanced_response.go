package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStackEnhancedResponse Response Object
type DeleteStackEnhancedResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStackEnhancedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStackEnhancedResponse struct{}"
	}

	return strings.Join([]string{"DeleteStackEnhancedResponse", string(data)}, " ")
}
