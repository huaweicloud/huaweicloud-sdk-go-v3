package v1

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/def"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/cbh/v1/model"
	"net/http"
)

func GenReqDefForListCbhInstance() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/cbs/instance/list").
		WithResponse(new(model.ListCbhInstanceResponse)).
		WithContentType("application/json")

	requestDef := reqDefBuilder.Build()
	return requestDef
}
