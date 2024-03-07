package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionCheckInDto struct {

	// 主对象ID。
	MasterId string `json:"masterId"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelVersionCheckInDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionCheckInDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionCheckInDto", string(data)}, " ")
}
