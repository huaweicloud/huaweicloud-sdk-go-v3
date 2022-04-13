package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteProjectV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProjectV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProjectV4Response struct{}"
	}

	return strings.Join([]string{"DeleteProjectV4Response", string(data)}, " ")
}
