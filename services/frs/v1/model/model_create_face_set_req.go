package model

import (
	"encoding/json"

	"strings"
)

type CreateFaceSetReq struct {
	// 用户自定义数据，自定义字段不能以系统保留字段vector、bounding_box、external_image_id、face_id、create_time、_id、_all、_source等字段命名。 Json字符串不校验重复性，自定义字段的key值长度范围为[1,36]，string类型的value长度范围为[1,256]，具体参见[自定义字段](zh-cn_topic_0130807044.xml)。

	ExternalFields map[string]TypeInfo `json:"external_fields,omitempty"`
	// 人脸库名称。 建议人脸库的名称不要以下划线（_）开头，否则云监控服务会无法采集人脸数量。

	FaceSetName string `json:"face_set_name"`
	// 人脸库最大的容量，填写1万整数倍的数字，例如：30000。 默认为100000，最大值为100000，可通过创建新的人脸库进行扩容，每个用户可免费默认使用10个人脸库，每个人脸库容量为10万个人脸特征。如需扩容单个人脸库规模，请联系华为云客服确认扩容规模与价格。

	FaceSetCapacity *int32 `json:"face_set_capacity,omitempty"`
}

func (o CreateFaceSetReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateFaceSetReq struct{}"
	}

	return strings.Join([]string{"CreateFaceSetReq", string(data)}, " ")
}
