package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplySecurityTableAuthorityRequest Request Object
type ApplySecurityTableAuthorityRequest struct {
	Body *ApplyTableAuthorityRequestBody `json:"body,omitempty"`
}

func (o ApplySecurityTableAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplySecurityTableAuthorityRequest struct{}"
	}

	return strings.Join([]string{"ApplySecurityTableAuthorityRequest", string(data)}, " ")
}
