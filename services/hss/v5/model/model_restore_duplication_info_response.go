package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreDuplicationInfoResponse Response Object
type RestoreDuplicationInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreDuplicationInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDuplicationInfoResponse struct{}"
	}

	return strings.Join([]string{"RestoreDuplicationInfoResponse", string(data)}, " ")
}
