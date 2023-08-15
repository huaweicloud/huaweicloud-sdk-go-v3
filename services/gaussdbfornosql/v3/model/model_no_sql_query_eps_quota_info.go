package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoSqlQueryEpsQuotaInfo struct {

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 企业项目名称。
	EnterpriseProjectName string `json:"enterprise_project_name"`

	Quota *NoSqlEpsQuotaTotal `json:"quota"`

	Used *NoSqlEpsQuotaUsed `json:"used"`
}

func (o NoSqlQueryEpsQuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlQueryEpsQuotaInfo struct{}"
	}

	return strings.Join([]string{"NoSqlQueryEpsQuotaInfo", string(data)}, " ")
}
