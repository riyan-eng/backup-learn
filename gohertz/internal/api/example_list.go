package api

import (
	"context"
	"gohertz/infrastructure"
	"gohertz/internal/dto"
	"gohertz/util"

	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary      List
// @Tags       	 Example
// @Produce      json
// @Param        order		query   string	false  "desc/asc default(desc)"
// @Param        search		query   string	false  "search"
// @Param        page		query   int		false  "page"
// @Param        per_page	query   int		false  "per_page"
// @Router       /example/ [get]
// @Security ApiKeyAuth
func (m *ServiceServer) ExampleList(ctx context.Context, c *app.RequestContext) {
	queryParam := new(dto.PaginationReq).Init()
	if err := c.BindQuery(&queryParam); err != nil {
		util.NewResponse(c).Error(err, "", 400)
		return
	}
	pageMeta := util.NewPagination().GetPageMeta(&queryParam.Page, &queryParam.Limit)

	data, countRow, err := m.exampleService.List(&ctx, pageMeta.Limit, pageMeta.Offset)
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	meta := util.PaginationMeta{
		Page:       pageMeta.Page,
		Limit:      pageMeta.Limit,
		CountRows:  countRow,
		CountPages: util.NewPagination().GetCountPages(countRow, pageMeta.Limit),
	}
	util.NewResponse(c).Success(data, meta, infrastructure.Localize("OK_READ"))
}
