package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeOrderRequest Request Object
type ChangeOrderRequest struct {
	Body *CbcOrderChange `json:"body,omitempty"`
}

func (o ChangeOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOrderRequest struct{}"
	}

	return strings.Join([]string{"ChangeOrderRequest", string(data)}, " ")
}
