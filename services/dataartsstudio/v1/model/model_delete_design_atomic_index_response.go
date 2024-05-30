package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesignAtomicIndexResponse Response Object
type DeleteDesignAtomicIndexResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDesignAtomicIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignAtomicIndexResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesignAtomicIndexResponse", string(data)}, " ")
}
