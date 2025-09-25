package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWebTamperRaspPathResponse Response Object
type ShowWebTamperRaspPathResponse struct {

	// **参数解释**: 动态网页防篡改的Tomcat bin目录 **取值范围**: 字符长度0-512位
	RaspPath       *string `json:"rasp_path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWebTamperRaspPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebTamperRaspPathResponse struct{}"
	}

	return strings.Join([]string{"ShowWebTamperRaspPathResponse", string(data)}, " ")
}
