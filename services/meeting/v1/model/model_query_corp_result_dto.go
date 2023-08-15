package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCorpResultDto struct {
	BasicInfo *QueryCorpBasicResultDto `json:"basicInfo,omitempty"`

	AdminInfo *QueryAdminResultDto `json:"adminInfo,omitempty"`

	ResInfo *QueryCorpResResultDto `json:"resInfo,omitempty"`

	GroupDTO *OrgGroupDto `json:"groupDTO,omitempty"`

	// 企业id。
	Id *string `json:"id,omitempty"`
}

func (o QueryCorpResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpResultDto", string(data)}, " ")
}
