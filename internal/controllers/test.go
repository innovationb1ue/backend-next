package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gopkg.in/guregu/null.v3"

	"github.com/penguin-statistics/backend-next/internal/models"
	"github.com/penguin-statistics/backend-next/internal/server"
	"github.com/penguin-statistics/backend-next/internal/service"
)

type TestController struct {
	fx.In
	DropMatrixService    *service.DropMatrixService
	DropInfoService      *service.DropInfoService
	PatternMatrixService *service.PatternMatrixService
	TrendService         *service.TrendService
}

func RegisterTestController(v3 *server.V3, c TestController) {
	v3.Get("/refresh/matrix/:server", c.RefreshAllDropMatrixElements)
	v3.Get("/global/matrix/:server", c.GetGlobalDropMatrix)
	v3.Get("/personal/matrix/:server/:accountId", c.GetPersonalDropMatrix)

	v3.Get("/advanced/matrix", c.AdvancedDropMatrixQuery)
	v3.Get("/advanced/trend", c.AdvancedTrendQuery)

	v3.Get("/refresh/pattern/:server", c.RefreshAllPatternMatrixElements)
	v3.Get("/global/pattern/:server", c.GetGlobalPatternMatrix)
	v3.Get("/personal/pattern/:server/:accountId", c.GetPersonalPatternMatrix)

	v3.Get("/refresh/trend/:server", c.RefreshAllTrendElements)
	v3.Get("/global/trend/:server", c.GetTrend)
}

func (c *TestController) RefreshAllDropMatrixElements(ctx *fiber.Ctx) error {
	server := ctx.Params("server")
	return c.DropMatrixService.RefreshAllDropMatrixElements(ctx, server)
}

func (c *TestController) GetGlobalDropMatrix(ctx *fiber.Ctx) error {
	server := ctx.Params("server")
	accountId := null.NewInt(0, false)
	globalDropMatrix, err := c.DropMatrixService.GetMaxAccumulableDropMatrixResults(ctx, server, &accountId)
	if err != nil {
		return err
	}
	return ctx.JSON(globalDropMatrix)
}

func (c *TestController) GetPersonalDropMatrix(ctx *fiber.Ctx) error {
	server := ctx.Params("server")
	accountIdStr := ctx.Params("accountId")
	accountIdNum, err := strconv.Atoi(accountIdStr)
	if err != nil {
		return err
	}
	accountIdNull := null.IntFrom(int64(accountIdNum))
	personalDropMatrix, err := c.DropMatrixService.GetMaxAccumulableDropMatrixResults(ctx, server, &accountIdNull)
	if err != nil {
		return err
	}
	return ctx.JSON(personalDropMatrix)
}

func (c *TestController) AdvancedDropMatrixQuery(ctx *fiber.Ctx) error {
	server := "CN"
	accountId := null.NewInt(75681, true)
	stageId := 383
	// itemIds := []int{7, 1}
	startTime := time.Unix(640966400, 0)
	endTime := time.Unix(11641052800, 0)
	timeRange := &models.TimeRange{StartTime: &startTime, EndTime: &endTime}
	elements, err := c.DropMatrixService.CalcDropMatrixForTimeRanges(ctx, server, []*models.TimeRange{timeRange}, []int{stageId}, nil, &accountId)
	if err != nil {
		return err
	}
	return ctx.JSON(elements)
}

func (c *TestController) RefreshAllPatternMatrixElements(ctx *fiber.Ctx) error {
	server := ctx.Params("server")
	return c.PatternMatrixService.RefreshAllPatternMatrixElements(ctx, server)
}

func (c *TestController) GetGlobalPatternMatrix(ctx *fiber.Ctx) error {
	server := ctx.Params("server")
	accountId := null.NewInt(0, false)
	globalPatternMatrix, err := c.PatternMatrixService.GetLatestPatternMatrixResults(ctx, server, &accountId)
	if err != nil {
		return err
	}
	return ctx.JSON(globalPatternMatrix)
}

func (c *TestController) GetPersonalPatternMatrix(ctx *fiber.Ctx) error {
	server := ctx.Params("server")
	accountIdStr := ctx.Params("accountId")
	accountIdNum, err := strconv.Atoi(accountIdStr)
	if err != nil {
		return err
	}
	accountIdNull := null.IntFrom(int64(accountIdNum))
	personalPatternMatrix, err := c.PatternMatrixService.GetLatestPatternMatrixResults(ctx, server, &accountIdNull)
	if err != nil {
		return err
	}
	return ctx.JSON(personalPatternMatrix)
}

func (c *TestController) RefreshAllTrendElements(ctx *fiber.Ctx) error {
	server := ctx.Params("server")
	return c.TrendService.RefreshTrendElements(ctx, server)
}

func (c *TestController) GetTrend(ctx *fiber.Ctx) error {
	server := ctx.Params("server")
	trend, err := c.TrendService.GetTrendResults(ctx, server)
	if err != nil {
		return err
	}
	return ctx.JSON(trend)
}

func (c *TestController) AdvancedTrendQuery(ctx *fiber.Ctx) error {
	server := "CN"
	accountId := null.NewInt(75681, false)
	stageId := 18
	itemIds := []int{7, 1}
	startTime := time.Unix(1640966400, 0)
	// endTime := time.Unix(1641052800, 0)
	// timeRange := &models.TimeRange{ StartTime: &startTime, EndTime: &endTime }
	elements, err := c.TrendService.CalcTrend(ctx, server, &startTime, 6, 3, []int{stageId}, itemIds, &accountId)
	if err != nil {
		return err
	}
	return ctx.JSON(elements)
}
