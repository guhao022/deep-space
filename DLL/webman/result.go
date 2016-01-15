package webman

type ERR_CODE int

const (
	CODE_SUCCESS ERR_CODE = iota + 2100		//
	CODE_ILLEGAL_DATA

	CODE_INSERT_ERR ERR_CODE = iota + 4100
	CODE_DEL_ERR
	CODE_UPDATE_ERR
	CODE_SELECT_ERR
)

type Res struct {
	Recode ERR_CODE    	`json:"recode"`
	Msg  string   		`json:"msg"`
	Data interface{} 	`json:"data"`
}


