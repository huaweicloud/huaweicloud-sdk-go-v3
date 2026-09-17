package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIsAdminUserNewRequest Request Object
type ShowIsAdminUserNewRequest struct {
}

func (o ShowIsAdminUserNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIsAdminUserNewRequest struct{}"
	}

	return strings.Join([]string{"ShowIsAdminUserNewRequest", string(data)}, " ")
}
