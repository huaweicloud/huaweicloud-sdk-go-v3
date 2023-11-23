package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLtsConfigRequest Request Object
type DeleteLtsConfigRequest struct {
	Body *DeleteLtsConfigRequestBody `json:"body,omitempty"`
}

func (o DeleteLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteLtsConfigRequest", string(data)}, " ")
}
