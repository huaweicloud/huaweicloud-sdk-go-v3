package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelFlinkJobResponse Response Object
type CancelFlinkJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelFlinkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelFlinkJobResponse struct{}"
	}

	return strings.Join([]string{"CancelFlinkJobResponse", string(data)}, " ")
}
