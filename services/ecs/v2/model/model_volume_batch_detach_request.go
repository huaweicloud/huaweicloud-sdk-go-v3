package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeBatchDetachRequest struct {
	Serverinfo *[]string `json:"serverinfo,omitempty"`
}

func (o VolumeBatchDetachRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeBatchDetachRequest struct{}"
	}

	return strings.Join([]string{"VolumeBatchDetachRequest", string(data)}, " ")
}
