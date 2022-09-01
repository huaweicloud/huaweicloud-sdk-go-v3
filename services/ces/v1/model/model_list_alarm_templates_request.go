package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmTemplatesRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 自定义告警模的ID，如：at1603330892378wkDm77y6B。
	AlarmTemplateId *string `json:"alarmTemplateId,omitempty" xml:"alarmTemplateId"`

	// 自定义告警模板选择的资源类型。即命名空间，如弹性云服务器的资源命名空间为：SYS.ECS；各服务命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 自定义告警模板选择的资源维度，如：弹性云服务器，则维度为instance_id，各资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Dname *string `json:"dname,omitempty" xml:"dname"`

	// 分页起始值，类型为integer，默认值为0。
	Start *string `json:"start,omitempty" xml:"start"`

	// 单次查询的条数限制，取值范围(0,100]，默认值为100， 用于限制结果数据条数。
	Limit *string `json:"limit,omitempty" xml:"limit"`
}

func (o ListAlarmTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplatesRequest", string(data)}, " ")
}
