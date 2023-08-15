package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSignatureFileRequest Request Object
type ShowSignatureFileRequest struct {

	// 营业执照ID
	FileId string `json:"file_id"`
}

func (o ShowSignatureFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSignatureFileRequest struct{}"
	}

	return strings.Join([]string{"ShowSignatureFileRequest", string(data)}, " ")
}
