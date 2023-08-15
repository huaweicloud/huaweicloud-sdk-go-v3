package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHealthmonitorRequest Request Object
type DeleteHealthmonitorRequest struct {

	// 健康检查id
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o DeleteHealthmonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthmonitorRequest", string(data)}, " ")
}
