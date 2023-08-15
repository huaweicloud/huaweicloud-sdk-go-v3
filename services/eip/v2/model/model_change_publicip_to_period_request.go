package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePublicipToPeriodRequest Request Object
type ChangePublicipToPeriodRequest struct {
	Body *ChangeToPeriodReq `json:"body,omitempty"`
}

func (o ChangePublicipToPeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePublicipToPeriodRequest struct{}"
	}

	return strings.Join([]string{"ChangePublicipToPeriodRequest", string(data)}, " ")
}
