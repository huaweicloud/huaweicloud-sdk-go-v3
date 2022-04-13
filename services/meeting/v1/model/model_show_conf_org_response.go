package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowConfOrgResponse struct {
	// 企业ID

	OrgID          *string `json:"orgID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConfOrgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfOrgResponse struct{}"
	}

	return strings.Join([]string{"ShowConfOrgResponse", string(data)}, " ")
}
