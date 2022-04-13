package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteApiGroupV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiGroupV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiGroupV2Response struct{}"
	}

	return strings.Join([]string{"DeleteApiGroupV2Response", string(data)}, " ")
}
