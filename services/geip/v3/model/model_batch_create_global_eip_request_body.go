package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateGlobalEipRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	GlobalEip *BatchCreateGlobalEipRequestBodyGlobalEip `json:"global_eip"`
}

func (o BatchCreateGlobalEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEipRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEipRequestBody", string(data)}, " ")
}
