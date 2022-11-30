package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type InstallApkRequest struct {
	Body *InstallApkRequestBody `json:"body,omitempty"`
}

func (o InstallApkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallApkRequest struct{}"
	}

	return strings.Join([]string{"InstallApkRequest", string(data)}, " ")
}
