package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFailureJobResponse Response Object
type DeleteFailureJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFailureJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFailureJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteFailureJobResponse", string(data)}, " ")
}
