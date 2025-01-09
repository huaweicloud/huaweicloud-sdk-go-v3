package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchInstallAppsRequest Request Object
type BatchInstallAppsRequest struct {
	Body *BatchAutoInstallAppsReq `json:"body,omitempty"`
}

func (o BatchInstallAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInstallAppsRequest struct{}"
	}

	return strings.Join([]string{"BatchInstallAppsRequest", string(data)}, " ")
}
