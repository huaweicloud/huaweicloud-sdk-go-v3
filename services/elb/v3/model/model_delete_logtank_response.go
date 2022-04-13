package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteLogtankResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogtankResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogtankResponse", string(data)}, " ")
}
