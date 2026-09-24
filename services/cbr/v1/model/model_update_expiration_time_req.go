package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateExpirationTimeReq struct {

	// 预期过期日期，格式：YYYY-MM-DD。
	ExpectExpirationDate string `json:"expect_expiration_date"`

	// 用户所在时区，格式形如 UTC+08:00
	TimeZone string `json:"time_zone"`
}

func (o UpdateExpirationTimeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExpirationTimeReq struct{}"
	}

	return strings.Join([]string{"UpdateExpirationTimeReq", string(data)}, " ")
}
