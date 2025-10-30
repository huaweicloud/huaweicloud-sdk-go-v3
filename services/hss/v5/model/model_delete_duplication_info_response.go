package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDuplicationInfoResponse Response Object
type DeleteDuplicationInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDuplicationInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDuplicationInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteDuplicationInfoResponse", string(data)}, " ")
}
