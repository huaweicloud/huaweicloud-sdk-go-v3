package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RemoveIngressEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveIngressEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveIngressEipV2Response struct{}"
	}

	return strings.Join([]string{"RemoveIngressEipV2Response", string(data)}, " ")
}
