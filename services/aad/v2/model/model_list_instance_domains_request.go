package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDomainsRequest Request Object
type ListInstanceDomainsRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`
}

func (o ListInstanceDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceDomainsRequest", string(data)}, " ")
}
