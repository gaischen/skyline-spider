package api_server


type Result struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Result  string `json:"result"`
}
