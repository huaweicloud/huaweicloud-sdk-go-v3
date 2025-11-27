package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceViewsRequest Request Object
type UpdateResourceViewsRequest struct {

	// 视图ID，长度1~32之间。
	Id string `json:"id"`

	Body *UpdateResourceViewsRequestBody `json:"body,omitempty"`
}

func (o UpdateResourceViewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceViewsRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceViewsRequest", string(data)}, " ")
}
