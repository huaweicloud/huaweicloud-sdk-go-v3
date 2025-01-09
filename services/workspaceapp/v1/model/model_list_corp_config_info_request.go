package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCorpConfigInfoRequest Request Object
type ListCorpConfigInfoRequest struct {
	Body *ListConfigInfoReq `json:"body,omitempty"`
}

func (o ListCorpConfigInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCorpConfigInfoRequest struct{}"
	}

	return strings.Join([]string{"ListCorpConfigInfoRequest", string(data)}, " ")
}
