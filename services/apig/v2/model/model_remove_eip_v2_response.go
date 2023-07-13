package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveEipV2Response Response Object
type RemoveEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveEipV2Response struct{}"
	}

	return strings.Join([]string{"RemoveEipV2Response", string(data)}, " ")
}
