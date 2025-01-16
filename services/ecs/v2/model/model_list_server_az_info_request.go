package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerAzInfoRequest Request Object
type ListServerAzInfoRequest struct {
}

func (o ListServerAzInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerAzInfoRequest struct{}"
	}

	return strings.Join([]string{"ListServerAzInfoRequest", string(data)}, " ")
}
