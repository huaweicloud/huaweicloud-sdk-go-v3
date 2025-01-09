package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthorizationMailRecordResponse Response Object
type ListAuthorizationMailRecordResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 授权邮件信息。
	Records        *[]AuthorizationMail `json:"records,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAuthorizationMailRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizationMailRecordResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorizationMailRecordResponse", string(data)}, " ")
}
