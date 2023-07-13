package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIterationV4Response Response Object
type DeleteIterationV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIterationV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIterationV4Response struct{}"
	}

	return strings.Join([]string{"DeleteIterationV4Response", string(data)}, " ")
}
