package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveDesignQualityInfosResponse Response Object
type RemoveDesignQualityInfosResponse struct {
	Data           *RemoveDesignQualityInfosResultData `json:"data,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o RemoveDesignQualityInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveDesignQualityInfosResponse struct{}"
	}

	return strings.Join([]string{"RemoveDesignQualityInfosResponse", string(data)}, " ")
}
