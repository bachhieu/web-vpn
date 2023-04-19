package utils

type ResCtrl struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResSuccess(status int, mess string, data interface{}) ResCtrl {
	return ResCtrl{status, mess, data}
}

func ResFail(status int, mess string, data string) ResCtrl {
	return ResCtrl{status, mess, data}
}

type ResData struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
