package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePoolRequest struct {

	// 后端云服务器组id
	PoolId string `json:"pool_id" xml:"pool_id"`

	Body *UpdatePoolRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdatePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolRequest struct{}"
	}

	return strings.Join([]string{"UpdatePoolRequest", string(data)}, " ")
}
