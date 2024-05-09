package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateCssConnectionRequest Request Object
type ValidateCssConnectionRequest struct {
	Body *CreateCssClusterReq `json:"body,omitempty"`
}

func (o ValidateCssConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateCssConnectionRequest struct{}"
	}

	return strings.Join([]string{"ValidateCssConnectionRequest", string(data)}, " ")
}
