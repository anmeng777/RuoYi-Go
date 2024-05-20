// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package ryserver

import (
	"RuoYi-Go/internal/handlers"
	"RuoYi-Go/internal/middleware"
	"github.com/kataras/iris/v12"
)

func StartServer(s *iris.Application) {
	s.Use(middleware.MiddlewareHandler)

	s.Get("/captchaImage", handler.CaptchaImage)
	s.Get("/getInfo", handler.GetInfo)

	s.Post("/login", handler.Login)
	s.Post("/logout", handler.Login)
}
