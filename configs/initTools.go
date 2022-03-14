package configs

import "url-shortener/utils/httpRouter"

func InitTools() httpRouter.Router {
	router := httpRouter.NewMuxRouter()

	return router
}
