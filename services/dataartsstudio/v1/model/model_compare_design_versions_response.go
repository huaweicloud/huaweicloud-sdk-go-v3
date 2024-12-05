package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareDesignVersionsResponse Response Object
type CompareDesignVersionsResponse struct {
	Data           *CompareDesignVersionsResultData `json:"data,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CompareDesignVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareDesignVersionsResponse struct{}"
	}

	return strings.Join([]string{"CompareDesignVersionsResponse", string(data)}, " ")
}
