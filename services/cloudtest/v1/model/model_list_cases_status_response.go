package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCasesStatusResponse Response Object
type ListCasesStatusResponse struct {
	Error *CommonResponseErrorOfobject `json:"error,omitempty"`

	Result *QueryCasesStatusResponseV2 `json:"result,omitempty"`

	// 状态值，如success、error
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCasesStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCasesStatusResponse struct{}"
	}

	return strings.Join([]string{"ListCasesStatusResponse", string(data)}, " ")
}
