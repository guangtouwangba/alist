package server

import (
	"fmt"
	"github.com/Xhofe/alist/drivers"
	"github.com/Xhofe/alist/model"
	"github.com/gofiber/fiber/v2"
)

func GetAccounts(ctx *fiber.Ctx) error {
	return SuccessResp(ctx, model.GetAccounts())
}

func SaveAccount(ctx *fiber.Ctx) error {
	var req model.Account
	if err := ctx.BodyParser(&req); err != nil {
		return ErrorResp(ctx, err, 400)
	}
	if err := validate.Struct(req); err != nil {
		return ErrorResp(ctx, err, 400)
	}
	driver, ok := drivers.GetDriver(req.Type)
	if !ok {
		return ErrorResp(ctx, fmt.Errorf("no [%s] driver", req.Type), 400)
	}
	if err := model.SaveAccount(req); err != nil {
		return ErrorResp(ctx, err, 500)
	} else {
		driver.Save(req)
		return SuccessResp(ctx)
	}
}

func DeleteAccount(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	if err := model.DeleteAccount(name); err != nil {
		return ErrorResp(ctx, err, 500)
	}
	return SuccessResp(ctx)
}