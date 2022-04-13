package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteApplicationV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationV4Response struct{}"
	}

	return strings.Join([]string{"DeleteApplicationV4Response", string(data)}, " ")
}
