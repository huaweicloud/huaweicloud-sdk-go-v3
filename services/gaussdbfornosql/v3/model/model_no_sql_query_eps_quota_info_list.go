package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoSqlQueryEpsQuotaInfoList struct {

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	Quota *NoSqlEpsQuota `json:"quota,omitempty"`

	Used *NoSqlQueryEpsQuotaUsed `json:"used,omitempty"`
}

func (o NoSqlQueryEpsQuotaInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlQueryEpsQuotaInfoList struct{}"
	}

	return strings.Join([]string{"NoSqlQueryEpsQuotaInfoList", string(data)}, " ")
}
