package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatastoresRequest Request Object
type ShowDatastoresRequest struct {
}

func (o ShowDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ShowDatastoresRequest", string(data)}, " ")
}
