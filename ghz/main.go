package main

import (
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	"github.com/golang/protobuf/proto"
	"github.com/infiniteprimes/portal-authzed/auth/pb"
	"log"
	"os"
	"time"
)

func main() {

	testingUpdateSchema()

}

const schema = `
definition user {}
`

func testingUpdateSchema() {
	var item = pb.UpdateSchemaRequest{
		Key:    "Hello",
		Schema: schema,
	}
	buf := proto.Buffer{}
	err := buf.EncodeMessage(&item)
	if err != nil {
		log.Fatal(err)
		return
	}

	report, err := runner.Run(
		// 基本配置 call host proto文件 data
		"authzed.v1.AuthzedSvc.UpdateSchema",
		"localhost:9192",
		//runner.WithDataFromJSON("/ghz.json"),
		runner.WithProtoFile("proto/authzed/authzed.proto", []string{}),
		//call Data
		runner.WithBinaryData(buf.Bytes()),

		//runner.WithBinaryDataFunc(dataFunc),

		runner.WithInsecure(true),
		//数量
		runner.WithTotalRequests(1000),

		//runner.WithStreamInterval(10*time.Second),
		// 并发参数
		//线性负荷
		runner.WithConcurrencySchedule(runner.ScheduleStep),
		runner.WithConcurrencyStep(100),
		runner.WithConcurrencyStart(5),
		runner.WithConcurrencyEnd(1000),
		runner.WithConcurrencyStepDuration(5*time.Second),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 指定输出路径
	file, err := os.Create("report.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	rp := printer.ReportPrinter{
		Out:    file,
		Report: report,
	}
	// 指定输出格式
	_ = rp.Print("html")
}
