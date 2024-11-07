package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseVersionRequest struct {

	// 目标版本
	TargetVersion string `json:"target_version"`
}

func (o DatabaseVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseVersionRequest struct{}"
	}

	return strings.Join([]string{"DatabaseVersionRequest", string(data)}, " ")
}
