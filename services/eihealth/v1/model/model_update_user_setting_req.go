package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateUserSettingReq struct {
	Operation *Operation `json:"operation"`

	Settings *UserSettingDto `json:"settings"`
}

func (o UpdateUserSettingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserSettingReq struct{}"
	}

	return strings.Join([]string{"UpdateUserSettingReq", string(data)}, " ")
}
