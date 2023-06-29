package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstancesV2Response Response Object
type DeleteInstancesV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstancesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesV2Response struct{}"
	}

	return strings.Join([]string{"DeleteInstancesV2Response", string(data)}, " ")
}
