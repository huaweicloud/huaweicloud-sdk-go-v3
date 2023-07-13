package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectQuotasResult struct {

	// 资源列表对象
	Resources []ResourceResult `json:"resources"`
}

func (o ProjectQuotasResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectQuotasResult struct{}"
	}

	return strings.Join([]string{"ProjectQuotasResult", string(data)}, " ")
}
