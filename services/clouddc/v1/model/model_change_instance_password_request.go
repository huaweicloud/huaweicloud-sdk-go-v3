package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstancePasswordRequest Request Object
type ChangeInstancePasswordRequest struct {
	Body *ChangeInstancePasswordOptions `json:"body,omitempty"`
}

func (o ChangeInstancePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstancePasswordRequest struct{}"
	}

	return strings.Join([]string{"ChangeInstancePasswordRequest", string(data)}, " ")
}
