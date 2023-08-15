package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoSqlModiflyEpsQuotasRequestBody struct {

	// 需要修改的企业项目配额信息列表。
	Quotas []NoSqlRequestEpsQuota `json:"quotas"`
}

func (o NoSqlModiflyEpsQuotasRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlModiflyEpsQuotasRequestBody struct{}"
	}

	return strings.Join([]string{"NoSqlModiflyEpsQuotasRequestBody", string(data)}, " ")
}
