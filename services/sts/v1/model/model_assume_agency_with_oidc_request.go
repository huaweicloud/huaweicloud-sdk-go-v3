package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyWithOidcRequest Request Object
type AssumeAgencyWithOidcRequest struct {
	Body *AssumeAgencyWithOidcReqBody `json:"body,omitempty"`
}

func (o AssumeAgencyWithOidcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyWithOidcRequest struct{}"
	}

	return strings.Join([]string{"AssumeAgencyWithOidcRequest", string(data)}, " ")
}
