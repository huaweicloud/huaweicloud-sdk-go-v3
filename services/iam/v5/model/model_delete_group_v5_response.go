package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupV5Response Response Object
type DeleteGroupV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGroupV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupV5Response struct{}"
	}

	return strings.Join([]string{"DeleteGroupV5Response", string(data)}, " ")
}
