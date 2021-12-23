package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type DetectLiveByFileRequestBody struct {
	// 本地视频文件。上传文件时，请求格式为multipart。 视频要求： • 视频文件大小不超过8MB，建议客户端压缩到200KB~2MB。 • 限制视频时长1～15秒。 • 建议帧率10fps～30fps。 • 封装格式：mp4、avi、flv、webm、asf、mov。 • 视频编码格式： h261、h263、h264、hevc、vc1、vp8、vp9、wmv3。
	VideoFile *def.FilePart `json:"video_file"`

	// 动作代码顺序列表，英文逗号（,）分隔。建议单动作，目前支持的动作有： • 1：左摇头 • 2：右摇头 • 3：点头 • 4：嘴部动作
	Actions *def.MultiPart `json:"actions"`

	// 该参数为动作时间数组拼接的字符串，数组的长度和actions的数量一致，每一项代表了对应次序动作的起始时间和结束时间，单位为距视频开始的毫秒数。
	ActionTime *def.MultiPart `json:"action_time,omitempty"`
}

func (o DetectLiveByFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByFileRequestBody struct{}"
	}

	return strings.Join([]string{"DetectLiveByFileRequestBody", string(data)}, " ")
}

func (o *DetectLiveByFileRequestBody) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o).Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]
		if m[jsonName] == nil && strings.Contains(jsonTag, "omitempty") {
			continue
		}
		field := v.FieldByName(utils.UnderscoreToCamel(jsonName))
		switch v.Field(i).Interface().(type) {
		case *def.FilePart:
			filePath := m[jsonName].(string)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(def.NewFilePart(file)))
		case *def.MultiPart:
			field.Set(reflect.ValueOf(def.NewMultiPart(m[jsonName])))
		}
	}
	return nil
}
