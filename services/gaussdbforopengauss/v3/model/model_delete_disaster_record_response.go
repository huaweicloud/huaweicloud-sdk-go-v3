package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDisasterRecordResponse Response Object
type DeleteDisasterRecordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDisasterRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecordResponse struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecordResponse", string(data)}, " ")
}
