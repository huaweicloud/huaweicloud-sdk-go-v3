package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinRequest Request Object
type UpdateRecycleBinRequest struct {
	Body *UpdateRecycleBinReq `json:"body,omitempty"`
}

func (o UpdateRecycleBinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinRequest", string(data)}, " ")
}
