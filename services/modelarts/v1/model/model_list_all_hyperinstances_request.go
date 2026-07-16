package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllHyperinstancesRequest Request Object
type ListAllHyperinstancesRequest struct {
}

func (o ListAllHyperinstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllHyperinstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAllHyperinstancesRequest", string(data)}, " ")
}
