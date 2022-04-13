package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDependencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDependencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDependencyResponse struct{}"
	}

	return strings.Join([]string{"DeleteDependencyResponse", string(data)}, " ")
}
