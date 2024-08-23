package bro_domains

type FleetLoginData struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponseData struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	TID     string         `json:"tid"`
	Data    LoginTokenData `json:"data"`
}

type LoginTokenData struct {
	FmsSessionID       string `json:"fms_session_id"`
	PasswordModifyNeed bool   `json:"password_modify_need"`
}

type TokenResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	TID     string            `json:"tid"`
	Data    TokenResponseData `json:"data"`
}

type TokenResponseData struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}

type ProfileResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Tid     string `json:"tid"`
	Data    struct {
		CompanyInfo struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Transportation bool   `json:"transportation"`
			Crowdsourcing  bool   `json:"crowdsourcing"`
		} `json:"company_info"`
		StaffInfo struct {
			ID      int    `json:"id"`
			Account string `json:"account"`
			Name    string `json:"name"`
		} `json:"staff_info"`
		MessageInfo struct {
			GrabOrderEnable bool `json:"grab_order_enable"`
		} `json:"message_info"`
		Modules []string `json:"modules"`
	} `json:"data"`
}

type Staff struct {
	CompanyId   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
	CompanyType int    `json:"company_type"`
}

type StaffResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TID     string `json:"tid"`
	Data    []Staff  `json:"data"`
}