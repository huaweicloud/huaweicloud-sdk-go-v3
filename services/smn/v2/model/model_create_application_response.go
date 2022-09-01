package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateApplicationResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// Application的唯一资源标识。
	ApplicationUrn *string `json:"application_urn,omitempty" xml:"application_urn"`

	// Application资源的ID。
	ApplicationId  *string `json:"application_id,omitempty" xml:"application_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationResponse", string(data)}, " ")
}
