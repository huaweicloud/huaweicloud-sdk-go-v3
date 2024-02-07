package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGlobalEipRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	GlobalEip *UpdateGlobalEipRequestBodyGlobalEip `json:"global_eip"`
}

func (o UpdateGlobalEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipRequestBody", string(data)}, " ")
}
