package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionUndoCheckOutDto struct {

	// 主对象ID。
	MasterId string `json:"masterId"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelVersionUndoCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionUndoCheckOutDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionUndoCheckOutDto", string(data)}, " ")
}
