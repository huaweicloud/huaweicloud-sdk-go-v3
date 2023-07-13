package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecordingsResponse Response Object
type DeleteRecordingsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRecordingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordingsResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecordingsResponse", string(data)}, " ")
}
