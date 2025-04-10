package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAntiVirusTaskResponse Response Object
type CreateAntiVirusTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAntiVirusTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiVirusTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAntiVirusTaskResponse", string(data)}, " ")
}
