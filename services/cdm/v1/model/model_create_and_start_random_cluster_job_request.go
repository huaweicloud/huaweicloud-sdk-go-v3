package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAndStartRandomClusterJobRequest struct {

	// 请求语言。
	XLanguage string `json:"X-Language" xml:"X-Language"`

	Body *CdmRandomCreateAndStartJobJsonReq `json:"body,omitempty" xml:"body"`
}

func (o CreateAndStartRandomClusterJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndStartRandomClusterJobRequest struct{}"
	}

	return strings.Join([]string{"CreateAndStartRandomClusterJobRequest", string(data)}, " ")
}
