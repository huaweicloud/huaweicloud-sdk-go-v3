package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePrivateBucketAccessRequest struct {
	// 加速域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`

	Body *UpdatePrivateBucketAccessBody `json:"body,omitempty"`
}

func (o UpdatePrivateBucketAccessRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePrivateBucketAccessRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateBucketAccessRequest", string(data)}, " ")
}
