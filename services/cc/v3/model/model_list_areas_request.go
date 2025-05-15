package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAreasRequest Request Object
type ListAreasRequest struct {
}

func (o ListAreasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAreasRequest struct{}"
	}

	return strings.Join([]string{"ListAreasRequest", string(data)}, " ")
}
