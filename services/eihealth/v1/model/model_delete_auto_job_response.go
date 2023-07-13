package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutoJobResponse Response Object
type DeleteAutoJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAutoJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutoJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteAutoJobResponse", string(data)}, " ")
}
