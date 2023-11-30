package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAtomicIndexByIdResponse Response Object
type ShowAtomicIndexByIdResponse struct {
	Data           *AtomicIndexVoDetailData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowAtomicIndexByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAtomicIndexByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowAtomicIndexByIdResponse", string(data)}, " ")
}
