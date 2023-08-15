package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountPublicIpResponse Response Object
type CountPublicIpResponse struct {

	// 弹性公网数量
	ElasticipSize  *int32 `json:"elasticip_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPublicIpResponse struct{}"
	}

	return strings.Join([]string{"CountPublicIpResponse", string(data)}, " ")
}
