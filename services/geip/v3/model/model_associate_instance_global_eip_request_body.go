package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateInstanceGlobalEipRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	GlobalEip *AssociateInstanceGlobalEipRequestBodyGlobalEip `json:"global_eip"`
}

func (o AssociateInstanceGlobalEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceGlobalEipRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateInstanceGlobalEipRequestBody", string(data)}, " ")
}
