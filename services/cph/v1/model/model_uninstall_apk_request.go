package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallApkRequest Request Object
type UninstallApkRequest struct {
	Body *UninstallApkRequestBody `json:"body,omitempty"`
}

func (o UninstallApkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallApkRequest struct{}"
	}

	return strings.Join([]string{"UninstallApkRequest", string(data)}, " ")
}
