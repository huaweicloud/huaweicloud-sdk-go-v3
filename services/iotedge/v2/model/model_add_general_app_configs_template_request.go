package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGeneralAppConfigsTemplateRequest Request Object
type AddGeneralAppConfigsTemplateRequest struct {
}

func (o AddGeneralAppConfigsTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGeneralAppConfigsTemplateRequest struct{}"
	}

	return strings.Join([]string{"AddGeneralAppConfigsTemplateRequest", string(data)}, " ")
}
