package core

import "fmt"

// logger.WithFields(logger.Fields{
//     "失败方法": utils.GetFuncName(),
// }).Fatal(core.FormatError(902, err).Error())

// logger.Info(core.FormatInfo(102))

var message = map[int]string{
	100: "成功",
	101: "配置文件读取成功",
	102: "服务启动开始",
	103: "服务关闭开始",
	104: "服务正在关闭",
	105: "服务中断信号收到",
	106: "服务启动成功",
	107: "服务关闭成功",
	108: "数据库连接启动",
	109: "数据库连接成功",
	110: "数据发送成功",

	299: "上游系统未知错误",

	300: "gRPC拨号失败",
	399: "下游系统未知错误",

	800: "数据库断开连接失败",
	801: "数据库连接失败",
	802: "数据库自动迁移失败",
	803: "所查询的数据不存在",

	900: "配置文件读取失败",
	901: "配置文件解析失败",
	902: "打开日志文件失败",
	903: "服务启动失败",
	904: "服务地址监听失败",
	998: "系统内部错误",
	999: "未知错误",
}

type Error struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Err  error  `json:"error"`
}

func (err *Error) Error() string {
	if err.Err != nil {
		return fmt.Sprintf("%s: %s", err.Desc, err.Err.Error())
	}
	return err.Desc
}

func FormatError(errorCode int, err error) *Error {
	var errObj = new(Error)
	errObj.Code = errorCode
	errObj.Desc = message[errorCode]
	errObj.Err = err
	return errObj
}

func FormatInfo(infoCode int) string {
	return message[infoCode]
}
