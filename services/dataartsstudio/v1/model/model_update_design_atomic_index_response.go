package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignAtomicIndexResponse Response Object
type UpdateDesignAtomicIndexResponse struct {
	Data           *AtomicIndexVoDetailData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o UpdateDesignAtomicIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignAtomicIndexResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesignAtomicIndexResponse", string(data)}, " ")
}
