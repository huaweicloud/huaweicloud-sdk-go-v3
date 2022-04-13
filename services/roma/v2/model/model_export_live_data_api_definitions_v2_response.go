package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Response Object
type ExportLiveDataApiDefinitionsV2Response struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportLiveDataApiDefinitionsV2Response) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportLiveDataApiDefinitionsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportLiveDataApiDefinitionsV2Response struct{}"
	}

	return strings.Join([]string{"ExportLiveDataApiDefinitionsV2Response", string(data)}, " ")
}
