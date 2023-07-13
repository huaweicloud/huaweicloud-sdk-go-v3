package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataJobResponse Response Object
type DeleteDataJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDataJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataJobResponse", string(data)}, " ")
}
