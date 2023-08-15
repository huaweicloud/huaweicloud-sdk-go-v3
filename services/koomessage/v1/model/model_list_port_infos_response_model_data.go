package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPortInfosResponseModelData struct {

	// 主键ID。
	Id *string `json:"id,omitempty"`

	// 创建时间，格式：yyyy-MM-ddTHH:mm:ssZ。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 服务号ID，在通道号列表显示为null。
	PubId *string `json:"pub_id,omitempty"`

	// 通道号。
	Port *string `json:"port,omitempty"`

	// 签名数组。  - 查询通道号列表时，该项为通道号签名列表  - 查询绑定关系时，该项为通道号绑定签名列表
	Sign *[]string `json:"sign,omitempty"`

	// 授权证明图片，key是上传的图片ID，value是图片对应的URL。
	AuthorizationFiles map[string]string `json:"authorization_files,omitempty"`

	// 服务号名称，查询通道号列表时该项为null。
	PubName *string `json:"pub_name,omitempty"`

	// 通道号类型。 - 1：普通 - 3：前缀号段 - 5：后缀号段
	PortType *int32 `json:"port_type,omitempty"`

	// 是否需要校验。  - 0：不校验 - 1：校验签名
	SignCheck *int32 `json:"sign_check,omitempty"`

	// 未绑定服务号时该项为null。
	Province *string `json:"province,omitempty"`

	// 是否绑定。  - 0: 未绑定 - 1: 绑定
	IsBind *int32 `json:"is_bind,omitempty"`

	// 绑定的服务号列表。  > 以JSON列表返回，格式： > [{\"pub_name\":\"服务号名称\",\"pub_reference\":\"服务号备注\"}]。
	PubList *[]PortSearchPubDetail `json:"pub_list,omitempty"`
}

func (o ListPortInfosResponseModelData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortInfosResponseModelData struct{}"
	}

	return strings.Join([]string{"ListPortInfosResponseModelData", string(data)}, " ")
}
