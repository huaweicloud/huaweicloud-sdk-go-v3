package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBucketsRequest Request Object
type AddBucketsRequest struct {

	// 资产类型
	Type *string `json:"type,omitempty"`

	Body *BucketsRequest `json:"body,omitempty"`
}

func (o AddBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBucketsRequest struct{}"
	}

	return strings.Join([]string{"AddBucketsRequest", string(data)}, " ")
}
