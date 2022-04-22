package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateYmlsResponse struct {

	// 返回值。
	Acknowledged *bool `json:"acknowledged,omitempty"`

	// 返回信息。
	ExternalMessage *string `json:"externalMessage,omitempty"`

	// 返回错误信息。
	HttpErrorResponse *string `json:"httpErrorResponse,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o UpdateYmlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateYmlsResponse struct{}"
	}

	return strings.Join([]string{"UpdateYmlsResponse", string(data)}, " ")
}
