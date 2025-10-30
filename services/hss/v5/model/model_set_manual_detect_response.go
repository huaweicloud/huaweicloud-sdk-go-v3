package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetManualDetectResponse Response Object
type SetManualDetectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetManualDetectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetManualDetectResponse struct{}"
	}

	return strings.Join([]string{"SetManualDetectResponse", string(data)}, " ")
}
