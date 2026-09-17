package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectsRequest Request Object
type ListEnterpriseProjectsRequest struct {
}

func (o ListEnterpriseProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectsRequest", string(data)}, " ")
}
