package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnterpriseProjectsResponse Response Object
type ShowEnterpriseProjectsResponse struct {

	// 集群企业项目列表
	Resources      *[]CdmClusterEnterpriseProject `json:"resources,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowEnterpriseProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectsResponse struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectsResponse", string(data)}, " ")
}
