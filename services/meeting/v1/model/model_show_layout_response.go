package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowLayoutResponse struct {

	// 结果码。 * 0：成功 * 非0：失败
	ReturnCode *int32 `json:"returnCode,omitempty"`

	// 结果描述。 * Success：成功 * 其他：失败原因
	ReturnDesc *string `json:"returnDesc,omitempty"`

	// 多画面布局。
	PicLayouts     *[]RestPicLayout `json:"picLayouts,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowLayoutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLayoutResponse struct{}"
	}

	return strings.Join([]string{"ShowLayoutResponse", string(data)}, " ")
}
