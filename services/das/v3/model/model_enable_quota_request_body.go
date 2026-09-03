package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableQuotaRequestBody 开通配额请求体
type EnableQuotaRequestBody struct {

	// 开通配额的数量
	OpenNum int32 `json:"open_num"`
}

func (o EnableQuotaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableQuotaRequestBody struct{}"
	}

	return strings.Join([]string{"EnableQuotaRequestBody", string(data)}, " ")
}
