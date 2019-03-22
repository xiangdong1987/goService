package main

import (
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"goService/model"
	"goService/myTool"
	"net/http"
	"os"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var svc FilterService
	//初始化词库
	dt := model.AcTrie{}
	dt.Dictionary = make(map[int32]int)
	path := myTool.GetCurrentPath()
	//fmt.Println(path)
	json, _ := myTool.ReadAll(path + "/model/dictionary/dictionary.json")
	var listDic map[string]string
	myTool.Jsondecode(json, &listDic)
	dt.InitDictionary(listDic)
	dt.Root = &model.AcNode{}
	dt.Root.Children = make([]*model.AcNode, dt.DicLength)
	for value, _ := range listDic {
		dt.AddWord(value)
	}
	//初始错误指针
	dt.InitFailPoint()
	svc = filterService{&dt}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}
	matchHandler := httptransport.NewServer(
		makeMatchEndpoint(svc),
		decodeMatchRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/match", matchHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
