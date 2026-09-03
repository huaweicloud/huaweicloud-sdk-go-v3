package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBinlogTaskResponse Response Object
type DeleteBinlogTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBinlogTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBinlogTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteBinlogTaskResponse", string(data)}, " ")
}
