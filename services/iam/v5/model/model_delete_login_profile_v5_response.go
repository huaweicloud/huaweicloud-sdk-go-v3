package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoginProfileV5Response Response Object
type DeleteLoginProfileV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoginProfileV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoginProfileV5Response struct{}"
	}

	return strings.Join([]string{"DeleteLoginProfileV5Response", string(data)}, " ")
}
