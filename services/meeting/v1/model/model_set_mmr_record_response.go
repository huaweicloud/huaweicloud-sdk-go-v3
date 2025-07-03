package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMmrRecordResponse Response Object
type SetMmrRecordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetMmrRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMmrRecordResponse struct{}"
	}

	return strings.Join([]string{"SetMmrRecordResponse", string(data)}, " ")
}
