package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVaultProjectTagRequest struct {
}

func (o ShowVaultProjectTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultProjectTagRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultProjectTagRequest", string(data)}, " ")
}
