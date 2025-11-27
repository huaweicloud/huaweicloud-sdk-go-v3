package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyPublicationResponse Response Object
type BatchModifyPublicationResponse struct {

	// 修改发布结果。
	Publications   *[]BatchOperateResponseInfo `json:"publications,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o BatchModifyPublicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyPublicationResponse struct{}"
	}

	return strings.Join([]string{"BatchModifyPublicationResponse", string(data)}, " ")
}
