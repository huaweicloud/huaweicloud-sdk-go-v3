package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListServiceSetDetailsRequest struct {

	// 服务组id
	SetId string `json:"set_id"`
}

func (o ListServiceSetDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSetDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceSetDetailsRequest", string(data)}, " ")
}
