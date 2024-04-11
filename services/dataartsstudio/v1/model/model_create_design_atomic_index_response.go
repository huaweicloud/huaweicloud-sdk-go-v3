package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesignAtomicIndexResponse Response Object
type CreateDesignAtomicIndexResponse struct {
	Data           *AtomicIndexVoDetailData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateDesignAtomicIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesignAtomicIndexResponse struct{}"
	}

	return strings.Join([]string{"CreateDesignAtomicIndexResponse", string(data)}, " ")
}
