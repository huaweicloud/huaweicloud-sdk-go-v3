package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFullSpeedRecordConfigResponse Response Object
type UpdateFullSpeedRecordConfigResponse struct {

	// 结果码
	ResultCode *string `json:"result_code,omitempty"`

	// 结果信息
	ResultDesc     *string `json:"result_desc,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFullSpeedRecordConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFullSpeedRecordConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateFullSpeedRecordConfigResponse", string(data)}, " ")
}
