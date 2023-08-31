package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEcnWithIegResponse Response Object
type DeleteEcnWithIegResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEcnWithIegResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEcnWithIegResponse struct{}"
	}

	return strings.Join([]string{"DeleteEcnWithIegResponse", string(data)}, " ")
}
