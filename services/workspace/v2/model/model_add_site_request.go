package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSiteRequest Request Object
type AddSiteRequest struct {
	Body *AddSiteReq `json:"body,omitempty"`
}

func (o AddSiteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSiteRequest struct{}"
	}

	return strings.Join([]string{"AddSiteRequest", string(data)}, " ")
}
