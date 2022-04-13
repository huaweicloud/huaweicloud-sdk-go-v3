package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteLiveDataApiV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"DeleteLiveDataApiV2Response", string(data)}, " ")
}
