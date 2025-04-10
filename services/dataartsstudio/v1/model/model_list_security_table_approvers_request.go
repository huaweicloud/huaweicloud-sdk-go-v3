package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityTableApproversRequest Request Object
type ListSecurityTableApproversRequest struct {
	Body *ListTableApproversRequestBody `json:"body,omitempty"`
}

func (o ListSecurityTableApproversRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityTableApproversRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityTableApproversRequest", string(data)}, " ")
}
