package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostRequestsRequest Request Object
type PostRequestsRequest struct {

	// qabot编号，UUID格式，如：303a0a00-c88a-43e3-aa2f-d5b8b9832b02。
	QabotId string `json:"qabot_id"`

	Body *PostOldRequestsReq `json:"body,omitempty"`
}

func (o PostRequestsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostRequestsRequest struct{}"
	}

	return strings.Join([]string{"PostRequestsRequest", string(data)}, " ")
}
