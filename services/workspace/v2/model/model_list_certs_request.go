package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertsRequest Request Object
type ListCertsRequest struct {
}

func (o ListCertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertsRequest struct{}"
	}

	return strings.Join([]string{"ListCertsRequest", string(data)}, " ")
}
