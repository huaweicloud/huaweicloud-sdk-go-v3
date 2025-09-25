package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadReportLogoResponse Response Object
type UploadReportLogoResponse struct {

	// **参数解释**： logo文件唯一标识ID **取值范围**： 字符长度0-256位
	LogoId         *string `json:"logo_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadReportLogoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadReportLogoResponse struct{}"
	}

	return strings.Join([]string{"UploadReportLogoResponse", string(data)}, " ")
}
