package v5

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/def"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/codecraft/v5/model"
	"net/http"
)

func GenReqDefForCreateCompetitionScore() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v5/competitions/score-infos").
		WithResponse(new(model.CreateCompetitionScoreResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListCompetitionWorks() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v5/competitions/works").
		WithResponse(new(model.ListCompetitionWorksResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("CompetitionId").
		WithJsonTag("competition_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StageId").
		WithJsonTag("stage_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ReadTime").
		WithJsonTag("read_time").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TimeUnit").
		WithJsonTag("time_unit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SortKey").
		WithJsonTag("sort_key").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SortDir").
		WithJsonTag("sort_dir").
		WithLocationType(def.Query))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRegisterCompetitionInfo() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v5/competitions/registrations").
		WithResponse(new(model.RegisterCompetitionInfoResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateCompetitionScore() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v5/competitions/scores").
		WithResponse(new(model.UpdateCompetitionScoreResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}
