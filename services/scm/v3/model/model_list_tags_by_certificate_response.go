package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsByCertificateResponse Response Object
type ListTagsByCertificateResponse struct {

	// 标签列表，key和value键值对的集合。
	Tags           *[]ScsResourceTag `json:"tags,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTagsByCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsByCertificateResponse struct{}"
	}

	return strings.Join([]string{"ListTagsByCertificateResponse", string(data)}, " ")
}
