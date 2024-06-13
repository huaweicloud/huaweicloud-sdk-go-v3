package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFreeDeclarationRequest Request Object
type ShowFreeDeclarationRequest struct {
}

func (o ShowFreeDeclarationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFreeDeclarationRequest struct{}"
	}

	return strings.Join([]string{"ShowFreeDeclarationRequest", string(data)}, " ")
}
