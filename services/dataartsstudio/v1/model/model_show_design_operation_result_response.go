package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesignOperationResultResponse Response Object
type ShowDesignOperationResultResponse struct {
	Data           *CheckDimensionStatusResultData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowDesignOperationResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesignOperationResultResponse struct{}"
	}

	return strings.Join([]string{"ShowDesignOperationResultResponse", string(data)}, " ")
}
