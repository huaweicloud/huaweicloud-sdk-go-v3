package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResEnterprisesRequest Request Object
type ListResEnterprisesRequest struct {
}

func (o ListResEnterprisesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResEnterprisesRequest struct{}"
	}

	return strings.Join([]string{"ListResEnterprisesRequest", string(data)}, " ")
}
