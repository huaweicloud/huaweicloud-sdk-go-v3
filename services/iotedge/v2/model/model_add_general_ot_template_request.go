package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddGeneralOtTemplateRequest struct {
}

func (o AddGeneralOtTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGeneralOtTemplateRequest struct{}"
	}

	return strings.Join([]string{"AddGeneralOtTemplateRequest", string(data)}, " ")
}
