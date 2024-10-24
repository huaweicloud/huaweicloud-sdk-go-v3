package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelVersionResponse Response Object
type DeleteModelVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteModelVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteModelVersionResponse", string(data)}, " ")
}
