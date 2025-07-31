package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAntiVirusPaidTaskResponse Response Object
type CreateAntiVirusPaidTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAntiVirusPaidTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiVirusPaidTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAntiVirusPaidTaskResponse", string(data)}, " ")
}
