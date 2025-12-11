package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAssociationRoutePolicyResponse Response Object
type ChangeAssociationRoutePolicyResponse struct {
	Association *Association `json:"association,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeAssociationRoutePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAssociationRoutePolicyResponse struct{}"
	}

	return strings.Join([]string{"ChangeAssociationRoutePolicyResponse", string(data)}, " ")
}
