package response

// Структуры для JSON описываются в файлике structs.go в каждом отдельном пакете.
// дальше делается кодогенерация при помощи библиотеки mailru/easyjson: easyjson -all structs.go,
// библиотека генерирует новый файл structs_easyjson.go в котором описаны методы MarshalJSON UnmarshalJSON для всех структур из structs.go
// Для работы с json необходимо использовать ТОЛЬКО эти методы.
type BaseResponse struct {
	Message string
	Status  int    `json:"status"`
	Error   string `json:"error,omitempty"`
}
