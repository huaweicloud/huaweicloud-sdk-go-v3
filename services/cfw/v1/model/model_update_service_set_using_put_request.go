package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateServiceSetUsingPutRequest struct {

	// 服务组id
	SetId string `json:"set_id"`

	Body *UpdateServiceSetUsingPutRequestBody `json:"body,omitempty"`
}

func (o UpdateServiceSetUsingPutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceSetUsingPutRequest struct{}"
	}

	return strings.Join([]string{"UpdateServiceSetUsingPutRequest", string(data)}, " ")
}
