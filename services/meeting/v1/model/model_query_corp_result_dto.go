package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCorpResultDto struct {
	BasicInfo *QueryCorpBasicResultDto `json:"basicInfo,omitempty" xml:"basicInfo"`

	AdminInfo *QueryAdminResultDto `json:"adminInfo,omitempty" xml:"adminInfo"`

	ResInfo *QueryCorpResResultDto `json:"resInfo,omitempty" xml:"resInfo"`

	GroupDTO *OrgGroupDto `json:"groupDTO,omitempty" xml:"groupDTO"`

	// 企业id
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o QueryCorpResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpResultDto", string(data)}, " ")
}
