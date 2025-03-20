package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessKeyV5Response Response Object
type DeleteAccessKeyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAccessKeyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessKeyV5Response struct{}"
	}

	return strings.Join([]string{"DeleteAccessKeyV5Response", string(data)}, " ")
}
