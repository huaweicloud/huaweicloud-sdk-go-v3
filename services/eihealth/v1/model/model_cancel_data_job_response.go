package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelDataJobResponse Response Object
type CancelDataJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelDataJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelDataJobResponse struct{}"
	}

	return strings.Join([]string{"CancelDataJobResponse", string(data)}, " ")
}
