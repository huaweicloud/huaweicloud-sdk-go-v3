package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PubInfoRequestBody PubRequestBody
type PubInfoRequestBody struct {

	// 服务号LOGO图片ID。  > 通过上传智能信息服务号图片资源接口上传LOGO图片后获得的图片resource_id。图片要求大于等于240*240且比例相同。
	LogoImg string `json:"logo_img"`

	// 服务号名称。  > 同一个企业下可以相同，不同企业下不能重复。
	PubName string `json:"pub_name"`

	// 服务号简介。
	PubAbstract string `json:"pub_abstract"`

	// 服务号备注。  > 同一个企业下，服务号名称相同时该项必须不同。
	PubRemark *string `json:"pub_remark,omitempty"`

	// 是否授权系统自动收集端口。   - 0：否 - 1：是
	AutoGetPort *int32 `json:"auto_get_port,omitempty"`

	// 自动收集端口使用的签名列表。  > auto_get_port为1时，该字段为必填，每个签名长度为2-18个字符，每个服务号签名不可以重复。
	SignsForAutoGetPort *[]string `json:"signs_for_auto_get_port,omitempty"`

	// 自动收集端口生效的地区列表。地区取值见《地区名称列表》。  > auto_get_port为1时，该字段有效。不填则默认全国，不允许传入重叠地区。
	AreasForAutoGetPort *[]string `json:"areas_for_auto_get_port,omitempty"`

	// 从事行业，默认取服务号所属商家的行业分类。  - 1：金融理财  - 2：社交通讯  - 3：影音娱乐  - 4：旅游出行  - 5：购物  - 6：本地生活  - 7：运动健康  - 8：教育培训  - 9：新闻阅读  - 10：运营商  - 11：其他
	Industry *int32 `json:"industry,omitempty"`

	// 授权证明图片ID，支持jpg、bmp、png和jpeg格式，全部图片总大小不超过4M，最多支持6张。  > 参数值为上传智能信息服务号图片资源API返回的resource_id。
	AuthorizationFiles []string `json:"authorization_files"`
}

func (o PubInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PubInfoRequestBody struct{}"
	}

	return strings.Join([]string{"PubInfoRequestBody", string(data)}, " ")
}
