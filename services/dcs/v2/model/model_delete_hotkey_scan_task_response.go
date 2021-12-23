package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteHotkeyScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHotkeyScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHotkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteHotkeyScanTaskResponse", string(data)}, " ")
}
