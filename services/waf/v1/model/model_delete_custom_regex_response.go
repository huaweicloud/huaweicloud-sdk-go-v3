package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomRegexResponse Response Object
type DeleteCustomRegexResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomRegexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRegexResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomRegexResponse", string(data)}, " ")
}
