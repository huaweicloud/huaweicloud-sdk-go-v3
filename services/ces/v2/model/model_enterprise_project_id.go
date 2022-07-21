package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业项目ID
type EnterpriseProjectId struct {
}

func (o EnterpriseProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectId struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectId", string(data)}, " ")
}
