package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinEnableRequest Request Object
type UpdateRecycleBinEnableRequest struct {
	Body *UpdateRecycleBinEnableRequestBody `json:"body,omitempty"`
}

func (o UpdateRecycleBinEnableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinEnableRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinEnableRequest", string(data)}, " ")
}
