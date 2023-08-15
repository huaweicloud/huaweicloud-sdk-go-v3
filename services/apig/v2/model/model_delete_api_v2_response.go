package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApiV2Response Response Object
type DeleteApiV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiV2Response struct{}"
	}

	return strings.Join([]string{"DeleteApiV2Response", string(data)}, " ")
}
