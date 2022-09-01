package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAppResponse struct {

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppResponse struct{}"
	}

	return strings.Join([]string{"CreateAppResponse", string(data)}, " ")
}
