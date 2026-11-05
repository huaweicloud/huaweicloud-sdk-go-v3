package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomRegexSetResponse Response Object
type DeleteCustomRegexSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomRegexSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRegexSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomRegexSetResponse", string(data)}, " ")
}
