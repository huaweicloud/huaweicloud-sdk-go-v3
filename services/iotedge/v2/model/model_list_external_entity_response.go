package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExternalEntityResponse Response Object
type ListExternalEntityResponse struct {

	// 外部实体列表
	Externals      *[]ExternalEntityRespDto `json:"externals,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListExternalEntityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalEntityResponse struct{}"
	}

	return strings.Join([]string{"ListExternalEntityResponse", string(data)}, " ")
}
