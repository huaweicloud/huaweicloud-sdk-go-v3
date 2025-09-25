package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSerialConsoleOptionsResponse Response Object
type UpdateSerialConsoleOptionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSerialConsoleOptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSerialConsoleOptionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateSerialConsoleOptionsResponse", string(data)}, " ")
}
