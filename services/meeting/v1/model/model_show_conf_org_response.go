package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowConfOrgResponse struct {

	// 企业ID
	OrgID          *string `json:"orgID,omitempty" xml:"orgID"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConfOrgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfOrgResponse struct{}"
	}

	return strings.Join([]string{"ShowConfOrgResponse", string(data)}, " ")
}
