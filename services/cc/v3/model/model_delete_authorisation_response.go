package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuthorisationResponse Response Object
type DeleteAuthorisationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAuthorisationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuthorisationResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuthorisationResponse", string(data)}, " ")
}
