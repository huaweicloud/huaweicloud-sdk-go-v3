package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMapTileRequest Request Object
type ShowMapTileRequest struct {

	// 缩放级别，取值范围[0~20]。
	Z int32 `json:"z"`

	// 缩放网格上的 X 坐标。 值必须在 [0, 2的z次方-1] 范围内。
	X int32 `json:"x"`

	// 缩放网格上的 Y 坐标。 值必须在 [0, 2的z次方-1] 范围内。
	Y int32 `json:"y"`

	// 签名消息头为：  Authorization: HMAC-SHA256 Clientid=xxxx,Expiry=xxxx,Signature=xxxx  HMAC-SHA256为固定签名算法，Clientid、Expiry、Signature的值从获取获取SAS token请求返回的消息体中获取，要求Clientid，Expiry，Signature同时存在。
	Authorization string `json:"Authorization"`
}

func (o ShowMapTileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMapTileRequest struct{}"
	}

	return strings.Join([]string{"ShowMapTileRequest", string(data)}, " ")
}
