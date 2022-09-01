package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCorpResponse struct {
	BasicInfo *QueryCorpBasicResultDto `json:"basicInfo,omitempty" xml:"basicInfo"`

	AdminInfo *QueryAdminResultDto `json:"adminInfo,omitempty" xml:"adminInfo"`

	ResInfo *QueryCorpResResultDto `json:"resInfo,omitempty" xml:"resInfo"`

	GroupDTO *OrgGroupDto `json:"groupDTO,omitempty" xml:"groupDTO"`

	// 企业id
	Id             *string `json:"id,omitempty" xml:"id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCorpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCorpResponse struct{}"
	}

	return strings.Join([]string{"ShowCorpResponse", string(data)}, " ")
}
