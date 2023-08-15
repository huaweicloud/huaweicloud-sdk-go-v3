package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveEngressEipV2Response Response Object
type RemoveEngressEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveEngressEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveEngressEipV2Response struct{}"
	}

	return strings.Join([]string{"RemoveEngressEipV2Response", string(data)}, " ")
}
