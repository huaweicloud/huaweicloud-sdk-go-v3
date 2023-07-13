package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessConfigRequest Request Object
type DeleteAccessConfigRequest struct {
	Body *DeleteAccessConfigRequestBody `json:"body,omitempty"`
}

func (o DeleteAccessConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteAccessConfigRequest", string(data)}, " ")
}
