package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelJobResponse Response Object
type CancelJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelJobResponse struct{}"
	}

	return strings.Join([]string{"CancelJobResponse", string(data)}, " ")
}
