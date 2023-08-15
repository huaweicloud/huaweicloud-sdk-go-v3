package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTagResponse Response Object
type BatchDeleteTagResponse struct {
	Body           *[]DeleteTagRsp `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o BatchDeleteTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagResponse", string(data)}, " ")
}
