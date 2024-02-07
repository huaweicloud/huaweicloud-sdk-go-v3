package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGlobalEipRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	GlobalEip *CreateGlobalEipRequestBodyGlobalEip `json:"global_eip"`
}

func (o CreateGlobalEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipRequestBody", string(data)}, " ")
}
