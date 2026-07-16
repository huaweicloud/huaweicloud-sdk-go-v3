package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllDevServersRequest Request Object
type ListAllDevServersRequest struct {
}

func (o ListAllDevServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllDevServersRequest struct{}"
	}

	return strings.Join([]string{"ListAllDevServersRequest", string(data)}, " ")
}
