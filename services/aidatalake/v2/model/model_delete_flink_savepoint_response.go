package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlinkSavepointResponse Response Object
type DeleteFlinkSavepointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFlinkSavepointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkSavepointResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlinkSavepointResponse", string(data)}, " ")
}
