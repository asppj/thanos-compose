package statsd

import (
	"fmt"
	"time"

	"gopkg.in/alexcesaro/statsd.v2"
)

var statsdClient *statsd.Client

func InitStatsdClient(prefix, addr string) (err error) {
	fmt.Printf("prefix:%s , addr:%s\n", prefix, addr)
	statsdClient, err = statsd.New(statsd.Prefix(prefix), statsd.Address(addr))
	return
}

func ReportSuccessResponse(statusCode int, path string) {
	bucket := fmt.Sprintf("resp.%d.%s.Success", statusCode, path)
	statsdClient.Increment(bucket)
}

func ReportNormalSuccessResponse() {
	bucket := fmt.Sprintf("resp.200.Success")
	statsdClient.Increment(bucket)
}

func ReportNormalErrorResponse() {
	bucket := fmt.Sprintf("resp.4xx.Err")
	statsdClient.Increment(bucket)
}

func ReportErrorResponse(statusCode int, path string) {
	bucket := fmt.Sprintf("resp.%d.%s", statusCode, path)
	statsdClient.Increment(bucket)
}

func ReportIgnoreErrorResponse(statusCode int, path string) {
	bucket := fmt.Sprintf("ignore.resp.%d.%s", statusCode, path)
	statsdClient.Increment(bucket)
}

func ReportCashNum(cashNum float64) {
	bucket := "cash"
	statsdClient.Count(bucket, cashNum)
}

func ReportCostTimeResponse(path string, method string, duration time.Duration) {
	bucket := fmt.Sprintf("api.%s.%s.cost", method, path)
	allApiBucket := "api.all.cost"
	timeMs := int(duration.Seconds() * 1000) //毫秒的时间数

	statsdClient.Timing(bucket, timeMs)
	statsdClient.Timing(allApiBucket, timeMs)
}

func ReportFeastResponse(statusCode int, moduleName string) {
	bucket := fmt.Sprintf("feast.%d.%s", statusCode, moduleName)
	statsdClient.Increment(bucket)
}

func ReportRedisError(source string) {
	statsdClient.Increment("redis.err." + source)
}

func ReportOutResponse(statusCode int, moduleName string) {
	bucket := fmt.Sprintf("httpcode.%d.%s", statusCode, moduleName)
	statsdClient.Increment(bucket)
}

func ReportWithdrawTaskletSuccessPassResponse() {
	bucket := "tasklet.machine_verify.200.pass"
	statsdClient.Increment(bucket)
}

func ReportWithdrawTaskletAdVerifySuccessPassResponse(msg string) {
	bucket := fmt.Sprintf("tasklet.200.ad.verify.%s.pass", msg)
	statsdClient.Increment(bucket)
}

func ReportWithdrawTaskletAdVerifyErrorResponse(msg string) {
	bucket := fmt.Sprintf("tasklet.400.ad.verify.%s", msg)
	statsdClient.Increment(bucket)
}

func ReportWithdrawTaskletAdVerifyUnPassResponse(msg string) {
	bucket := fmt.Sprintf("tasklet.200.ad.verify.%s.unpass", msg)
	statsdClient.Increment(bucket)
}

func ReportWithdrawTaskletErrorPassResponse() {
	bucket := "tasklet.machine_verify.400"
	statsdClient.Increment(bucket)
}

func ReportWithdrawTaskletSuccessUnpassPassResponse(unPassCode int) {
	bucket := fmt.Sprintf("tasklet.machine_verify.200.unpass.%d", unPassCode)
	statsdClient.Increment(bucket)
}

func SyncTridentData(status int) {
	bucket := fmt.Sprintf("tasklet.sync_goldeneye.%d", status)
	statsdClient.Increment(bucket)
}

func ReportRankingGetTop100Err() {
	bucket := fmt.Sprintf("model.ranking.gettop100.err")
	statsdClient.Increment(bucket)
}

func ReportRankingGetScoreErr() {
	bucket := fmt.Sprintf("model.ranking.getscore.err")
	statsdClient.Increment(bucket)
}

func ReportRankingGetRankErr() {
	bucket := fmt.Sprintf("model.ranking.getrank.err")
	statsdClient.Increment(bucket)
}

func ReportRankingGetMemberErr() {
	bucket := fmt.Sprintf("model.ranking.getmember.err")
	statsdClient.Increment(bucket)
}

func ReportRankingIncrMemberErr() {
	bucket := fmt.Sprintf("model.ranking.incr.err")
	statsdClient.Increment(bucket)
}

func ReportFeastWithdrawResponse(code int) {
	bucket := fmt.Sprintf("resp.feast.withdraw.%d", code)
	statsdClient.Increment(bucket)

}

func ReportRiskControlAction(action int) {
	bucket := fmt.Sprintf("riskcontrol.action.%d", action)
	statsdClient.Increment(bucket)
}

func ReportWithdrawAction(action int) {
	bucket := fmt.Sprintf("withdraw.action.%d", action)
	statsdClient.Increment(bucket)
}

// 上报 server_host
func ReportServerHost(host string) {
	statsdClient.Increment("req.server.host." + host)
}

// 上报客户端传来的 config_version 和服务器存的不一致
func ReportUserConfigVersionChanged() {
	statsdClient.Increment("user.config_version.changed")
}

// 上报接口配置错误
func ReportConfigError(name string) {
	statsdClient.Increment(name + ".config.err")
}

func CloseStatsdClient() {
	statsdClient.Close()
}
