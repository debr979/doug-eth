package errMsg

import "fmt"

type ResultMsg string

type WithStatus struct {
	Msg    string
	Status int
}

var Messenger Msg

type Msg struct {
	Status  int
	Message string
}

const (
	//success
	AddSuccess    ResultMsg = "AddSuccess"
	GetSuccess    ResultMsg = "GetSuccess"
	DeleteSuccess ResultMsg = "DeleteSuccess"
	ModifySuccess ResultMsg = "ModifySuccess"

	//faild
	ErrorParams        ResultMsg = "ErrorParams"
	WrongPassword      ResultMsg = "WrongPassword"
	NoData             ResultMsg = "NoData"
	NotMember          ResultMsg = "NotMember"
	AddFailed          ResultMsg = "AddFailed"
	GetFailed          ResultMsg = "GetFailed"
	ModifyFailed       ResultMsg = "ModifyFailed"
	DeleteFailed       ResultMsg = "DeleteFailed"
	Unauthorized       ResultMsg = "Unauthorized"
	ErrAccOrPwd        ResultMsg = "ErrAccOrPwd"
	PasswordExceedTime ResultMsg = "PasswordExceedTime"
	AddressBeUsed      ResultMsg = "AddressBeUsed"
	EmailBeUsed        ResultMsg = "EmailBeUsed"
	EmailNotBound      ResultMsg = "EmailNotBound"

	//server error
	SystemError                ResultMsg = `SystemError`
	DatabaseError              ResultMsg = `DatabaseError`
	MsgSettingError            ResultMsg = `MsgSettingError`
	CallOuterApiError          ResultMsg = "CallOuterApiError"
	ContractInteractionFailure ResultMsg = "ContractInteractionFailure"
)

var errorMsgDict = map[ResultMsg]WithStatus{
	`AddSuccess`:    {`add success`, 2000},
	`GetSuccess`:    {`get success`, 2000},
	`DeleteSuccess`: {`delete success`, 2000},
	`ModifySuccess`: {`modify success`, 2000},
	//failed
	`ErrorParams`:        {`error params`, 4000},
	`NoData`:             {`no data`, 4001},
	`NotMember`:          {`not member`, 4002},
	`AddFailed`:          {`add faild`, 4003},
	`GetFailed`:          {`get failed`, 4004},
	`ModifyFailed`:       {`modify failed`, 4005},
	`DeleteFailed`:       {`delete failed`, 4006},
	`Unauthorized`:       {`unauthorized`, 4007},
	`ErrAccOrPwd`:        {`wrong email or password`, 4008},
	`PasswordExceedTime`: {`password exceed time`, 4009},
	`AddressBeUsed`:      {`address be used`, 4010},
	`EmailBeUsed`:        {`email be used`, 4011},
	`EmailNotBound`:      {`email not bound`, 4012},
	//server error
	`SystemError`:                {`system error`, 5000},
	`DatabaseError`:              {`database error`, 5001},
	`MsgSettingError`:            {`msg setting error`, 5002},
	`CallOuterApiError`:          {`Call outer api failed`, 5003},
	`ContractInteractionFailure`: {`Contract Interaction Failure`, 5004},
}

/*
GetResultMsg ...
*/
func GetResultMsg(key ResultMsg, params []interface{}) WithStatus {
	val, ok := errorMsgDict[key]
	if !ok {
		return errorMsgDict[MsgSettingError]
	}

	if len(params) > 0 {
		val.Msg = fmt.Sprintf(val.Msg, params...)
	}

	return val
}

func SetResultMsg(msg string, status int) WithStatus {
	var result WithStatus
	result.Msg = msg
	result.Status = status

	return result
}

func (r *Msg) Get(key ResultMsg, params ...interface{}) *Msg {

	// Get Message and Status
	result := GetResultMsg(key, params)

	r.Message = result.Msg
	r.Status = result.Status

	return r
}

func (r *Msg) Set(msgString string, status int) *Msg {
	r.Message = msgString
	r.Status = status

	return r
}
