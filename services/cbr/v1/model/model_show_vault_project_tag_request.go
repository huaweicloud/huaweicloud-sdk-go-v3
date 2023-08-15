package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVaultProjectTagRequest Request Object
type ShowVaultProjectTagRequest struct {
}

func (o ShowVaultProjectTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultProjectTagRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultProjectTagRequest", string(data)}, " ")
}
