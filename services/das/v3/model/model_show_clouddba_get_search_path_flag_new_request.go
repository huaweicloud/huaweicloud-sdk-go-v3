package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClouddbaGetSearchPathFlagNewRequest Request Object
type ShowClouddbaGetSearchPathFlagNewRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`
}

func (o ShowClouddbaGetSearchPathFlagNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClouddbaGetSearchPathFlagNewRequest struct{}"
	}

	return strings.Join([]string{"ShowClouddbaGetSearchPathFlagNewRequest", string(data)}, " ")
}
