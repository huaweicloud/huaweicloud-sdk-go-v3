package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRetrievalRequestBody struct {

	// 公网域名。
	ZoneName *string `json:"zone_name,omitempty"`
}

func (o CreateRetrievalRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetrievalRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRetrievalRequestBody", string(data)}, " ")
}
