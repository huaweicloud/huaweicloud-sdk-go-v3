package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomAnswersResponse Response Object
type DeleteCustomAnswersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomAnswersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomAnswersResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomAnswersResponse", string(data)}, " ")
}
