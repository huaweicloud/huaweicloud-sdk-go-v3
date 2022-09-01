package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCaseCcEmailsResponse struct {
	CcEmailInfo *IncidentOrderCcEmailInfoV2 `json:"cc_email_info,omitempty" xml:"cc_email_info"`

	// 抄送邮箱信息
	McEmailInfos   *[]string `json:"mc_email_infos,omitempty" xml:"mc_email_infos"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCaseCcEmailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaseCcEmailsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseCcEmailsResponse", string(data)}, " ")
}
