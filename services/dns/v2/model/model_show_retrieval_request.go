package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRetrievalRequest Request Object
type ShowRetrievalRequest struct {

	// 公网域名。
	Name string `json:"name"`
}

func (o ShowRetrievalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRetrievalRequest struct{}"
	}

	return strings.Join([]string{"ShowRetrievalRequest", string(data)}, " ")
}
