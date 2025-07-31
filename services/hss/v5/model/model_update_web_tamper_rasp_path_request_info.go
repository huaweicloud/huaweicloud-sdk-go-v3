package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWebTamperRaspPathRequestInfo struct {

	// rasp path
	RaspPath *string `json:"rasp_path,omitempty"`
}

func (o UpdateWebTamperRaspPathRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperRaspPathRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperRaspPathRequestInfo", string(data)}, " ")
}
