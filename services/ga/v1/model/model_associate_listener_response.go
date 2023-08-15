package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateListenerResponse Response Object
type AssociateListenerResponse struct {

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateListenerResponse struct{}"
	}

	return strings.Join([]string{"AssociateListenerResponse", string(data)}, " ")
}
