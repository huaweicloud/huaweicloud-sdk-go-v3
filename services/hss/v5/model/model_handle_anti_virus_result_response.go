package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleAntiVirusResultResponse Response Object
type HandleAntiVirusResultResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HandleAntiVirusResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleAntiVirusResultResponse struct{}"
	}

	return strings.Join([]string{"HandleAntiVirusResultResponse", string(data)}, " ")
}
