package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyWithSamlRequest Request Object
type AssumeAgencyWithSamlRequest struct {
	Body *AssumeAgencyWithSamlReqBody `json:"body,omitempty"`
}

func (o AssumeAgencyWithSamlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyWithSamlRequest struct{}"
	}

	return strings.Join([]string{"AssumeAgencyWithSamlRequest", string(data)}, " ")
}
