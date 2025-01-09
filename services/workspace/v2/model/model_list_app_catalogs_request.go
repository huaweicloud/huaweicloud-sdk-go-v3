package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppCatalogsRequest Request Object
type ListAppCatalogsRequest struct {
}

func (o ListAppCatalogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppCatalogsRequest struct{}"
	}

	return strings.Join([]string{"ListAppCatalogsRequest", string(data)}, " ")
}
