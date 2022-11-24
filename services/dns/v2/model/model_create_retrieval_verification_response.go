package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRetrievalVerificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRetrievalVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetrievalVerificationResponse struct{}"
	}

	return strings.Join([]string{"CreateRetrievalVerificationResponse", string(data)}, " ")
}
