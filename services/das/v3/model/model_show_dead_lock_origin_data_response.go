package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockOriginDataResponse Response Object
type ShowDeadLockOriginDataResponse struct {

	// 原始数据
	OriginData     *string `json:"origin_data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDeadLockOriginDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockOriginDataResponse struct{}"
	}

	return strings.Join([]string{"ShowDeadLockOriginDataResponse", string(data)}, " ")
}
