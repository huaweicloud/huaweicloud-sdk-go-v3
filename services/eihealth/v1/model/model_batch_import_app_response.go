package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportAppResponse Response Object
type BatchImportAppResponse struct {
	Body           *[]ImportAppRsp `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o BatchImportAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportAppResponse struct{}"
	}

	return strings.Join([]string{"BatchImportAppResponse", string(data)}, " ")
}
