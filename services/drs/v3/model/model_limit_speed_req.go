package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type LimitSpeedReq struct {

	// 任务id
	JobId string `json:"job_id" xml:"job_id"`

	// 限速信息请求体
	SpeedLimit []SpeedLimitInfo `json:"speed_limit" xml:"speed_limit"`
}

func (o LimitSpeedReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LimitSpeedReq struct{}"
	}

	return strings.Join([]string{"LimitSpeedReq", string(data)}, " ")
}
