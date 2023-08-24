package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppConfigsTemplatesRequest Request Object
type AddAppConfigsTemplatesRequest struct {
	Body *CreateAppConfigsTemplatesReqDto `json:"body,omitempty"`
}

func (o AddAppConfigsTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppConfigsTemplatesRequest struct{}"
	}

	return strings.Join([]string{"AddAppConfigsTemplatesRequest", string(data)}, " ")
}
