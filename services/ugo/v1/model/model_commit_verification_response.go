package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitVerificationResponse Response Object
type CommitVerificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CommitVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitVerificationResponse struct{}"
	}

	return strings.Join([]string{"CommitVerificationResponse", string(data)}, " ")
}
