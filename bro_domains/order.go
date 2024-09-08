package bro_domains

type BaseInfo struct {
	// SerialNo        string `json:"serial_no"`
	SerialId        string `json:"serial_id"`
	Track           string `json:"track"`
	Area            string `json:"area"`
	CarType         int    `json:"car_type"`
	CarTypeText     string `json:"car_type_text"`
	RunningMileage  string `json:"running_mileage"`
	StopTime        int    `json:"stop_time"`
	FinishForTime   string `json:"finish_for_time"`
	PlanArrivedTime int    `json:"plan_arrived_time"`
	Price           int    `json:"price"`
	CreatedAt       uint   `json:"created_at"`
}

type SiteInfo struct {
	OrderNo           int     `json:"order_no"`
	StoreName         string  `json:"store_name"`
	Address           string  `json:"address"`
	RunningMileage    *string `json:"running_mileage,omitempty"`
	PlanArrivedTime   int     `json:"plan_arrived_time"`
	PlanDepartureTime *int    `json:"plan_departure_time,omitempty"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
}

type Order struct {
	BaseInfo BaseInfo   `json:"base_info"`
	SiteInfo []SiteInfo `json:"site_info"`
}

type OrderQuery struct {
	Area     int `json:"area"`
	CarType  int `json:"car_type"`
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

type OrderQueryAll struct {
	Area     string `json:"area"`
	CarType  string `json:"car_type"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}

type OrderSubmit struct {
	SerialId   string `json:"serial_id"`
	CheckToken string `json:"check_token"`
}

type OrderResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Tid     string            `json:"tid"`
	Data    OrderResponseData `json:"data"`
}

type OrderResponseData struct {
	Items      []Order                 `json:"items"`
	Pagination OrderResponsePagination `json:"pagination"`
}

type OrderResponsePagination struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	TotalCount  int `json:"totalCount"`
}

type DriverReq struct {
	DriverID        string `json:"driver_id"`
	Label           string `json:"label"`
	CarCompanyID    string `json:"car_company_id"`
	DriverCardState string `json:"driver_card_state"`
	AuditState      string `json:"audit_state"`
	Type            string `json:"type"`
	ExternalStaffID string `json:"external_staff_id"`
	PageSize        int    `json:"page_size"`
	PageNum         int    `json:"page_num"`
	TotalCount      int    `json:"totalCount"`
}

type DriverRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Tid     string `json:"tid"`
	Data    struct {
		Items []struct {
			ID                  int         `json:"id"`
			DriverID            int         `json:"driver_id"`
			ExternalStaffID     interface{} `json:"external_staff_id"`
			DriverCardEnd       interface{} `json:"driver_card_end"`
			DriverCardState     interface{} `json:"driver_card_state"`
			DriverCardStateText string      `json:"driver_card_state_text"`
			AuditState          int         `json:"audit_state"`
			AuditStateText      string      `json:"audit_state_text"`
			DriverName          string      `json:"driver_name"`
			Label               interface{} `json:"label"`
			LabelText           string      `json:"label_text"`
			Mobile              string      `json:"mobile"`
			Nation              int         `json:"nation"`
			NationText          string      `json:"nation_text"`
			Source              int         `json:"source"`
			CarCnt              int         `json:"car_cnt"`
			SourceText          string      `json:"source_text"`
			CardNum             string      `json:"card_num"`
			DriverCardNum       string      `json:"driver_card_num"`
			BlacklistStatus     int         `json:"blacklist_status"`
			BlacklistStatusText interface{} `json:"blacklist_status_text"`
		} `json:"items"`
		Pagination struct {
			CurrentPage int `json:"currentPage"`
			PerPage     int `json:"perPage"`
			TotalCount  int `json:"totalCount"`
		} `json:"pagination"`
	} `json:"data"`
}

type CarReq struct {
	CarNum       string `json:"car_num"`
	CarCompanyID string `json:"car_company_id"`
	CarState     string `json:"car_state"`
	Type         string `json:"type"`
	BizType      string `json:"biz_type"`
	Size         string `json:"size"`
	PageNum      int    `json:"page_num"`
	PageSize     int    `json:"page_size"`
}

type CarRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Tid     string `json:"tid"`
	Data    struct {
		Items []struct {
			ID             int `json:"id"`
			CompanyID      int `json:"company_id"`
			CarCompanyList []struct {
				CarCompanyID   int    `json:"car_company_id"`
				CarCompanyName string `json:"car_company_name"`
			} `json:"car_company_list"`
			Type           int         `json:"type"`
			TypeText       string      `json:"type_text"`
			BindType       int         `json:"bind_type"`
			Size           int         `json:"size"`
			SizeText       string      `json:"size_text"`
			Brand          int         `json:"brand"`
			BrandText      string      `json:"brand_text"`
			CarNum         string      `json:"car_num"`
			CarState       int         `json:"car_state"`
			CarStateText   string      `json:"car_state_text"`
			BizType        int         `json:"biz_type"`
			BizTypeText    string      `json:"biz_type_text"`
			InitMileage    int         `json:"init_mileage"`
			CurrentMileage int         `json:"current_mileage"`
			OilType        int         `json:"oil_type"`
			OilNum         int         `json:"oil_num"`
			BuyDate        interface{} `json:"buy_date"`
			CreateUID      int         `json:"create_uid"`
			CreatedAt      int         `json:"created_at"`
			UpdatedAt      int         `json:"updated_at"`
			Province       string      `json:"province"`
			City           string      `json:"city"`
			District       string      `json:"district"`
			PostalCode     string      `json:"postal_code"`
			ThCode         string      `json:"th_code"`
			Address        string      `json:"address"`
			CarProvince    string      `json:"car_province"`
			DriverNum      int         `json:"driver_num"`
			GpsInfo        interface{} `json:"gps_info"`
		} `json:"items"`
		Pagination struct {
			CurrentPage int `json:"currentPage"`
			PerPage     int `json:"perPage"`
			TotalCount  int `json:"totalCount"`
		} `json:"pagination"`
	} `json:"data"`
}

type OrderDetailResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Tid     string     `json:"tid"`
	Data    []SiteInfo `json:"data"`
}

type OrderRearangeReq struct {
	CarType  int     `json:"car_type"`
	CarID    int     `json:"car_id"`
	DriverID int     `json:"driver_id"`
	SerialID string  `json:"serial_id"`
	LineType *string `json:"line_type"`
	Province string  `json:"province"`
}
