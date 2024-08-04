package web

type Pagination struct {
	PageSize    int `json:"page_size"`
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	TotalItems  int `json:"total_items"`
}

type WebResponsePagination struct {
	Code       int         `json:"code"`
	Status     string      `json:"status"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}