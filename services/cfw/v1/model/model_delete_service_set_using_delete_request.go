package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteServiceSetUsingDeleteRequest struct {

	// 服务集合id
	SetId string `json:"set_id"`
}

func (o DeleteServiceSetUsingDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceSetUsingDeleteRequest struct{}"
	}

	return strings.Join([]string{"DeleteServiceSetUsingDeleteRequest", string(data)}, " ")
}
