package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDDosStatusResponse struct {

	// 防护状态，可选范围：   - normal：表示正常   - configging：表示设置中   - notConfig：表示未设置   - packetcleaning：表示清洗   - packetdropping：表示黑洞
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDDosStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDDosStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowDDosStatusResponse", string(data)}, " ")
}
