package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListEnterpriseProjectsResDetail struct {
	// 项目Id。

	ProjectId string `json:"projectId"`
}

func (o ListEnterpriseProjectsResDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectsResDetail struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectsResDetail", string(data)}, " ")
}
