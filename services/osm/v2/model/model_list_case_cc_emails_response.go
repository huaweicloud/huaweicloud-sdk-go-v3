package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCaseCcEmailsResponse struct {
	CcEmailInfo *IncidentOrderCcEmailInfoV2 `json:"cc_email_info,omitempty"`
	// 抄送邮箱信息

	McEmailInfos   *[]string `json:"mc_email_infos,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCaseCcEmailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaseCcEmailsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseCcEmailsResponse", string(data)}, " ")
}
