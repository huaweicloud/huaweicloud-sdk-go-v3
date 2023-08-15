package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQaPairDetailRequest Request Object
type ShowQaPairDetailRequest struct {

	// 调用智能客服服务标志。
	XServiceKey string `json:"x-service-key"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	// 语料Id
	QaPairId string `json:"qa_pair_id"`
}

func (o ShowQaPairDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQaPairDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowQaPairDetailRequest", string(data)}, " ")
}
