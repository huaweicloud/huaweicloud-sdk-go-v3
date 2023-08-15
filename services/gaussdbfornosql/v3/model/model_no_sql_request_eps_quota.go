package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoSqlRequestEpsQuota struct {

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Quota *NoSqlEpsQuotaRequestInfo `json:"quota"`
}

func (o NoSqlRequestEpsQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlRequestEpsQuota struct{}"
	}

	return strings.Join([]string{"NoSqlRequestEpsQuota", string(data)}, " ")
}
