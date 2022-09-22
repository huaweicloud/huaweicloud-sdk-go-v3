package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowConfOrgResponse struct {

	// SP管理员根据会议ID查询该会议归属的企业ID。
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
