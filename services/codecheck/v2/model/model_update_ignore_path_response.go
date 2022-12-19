package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateIgnorePathResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateIgnorePathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIgnorePathResponse struct{}"
	}

	return strings.Join([]string{"UpdateIgnorePathResponse", string(data)}, " ")
}
