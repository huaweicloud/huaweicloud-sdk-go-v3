package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCaseResponse Response Object
type UpdateCaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCaseResponse struct{}"
	}

	return strings.Join([]string{"UpdateCaseResponse", string(data)}, " ")
}
