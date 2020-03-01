package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/brianvoe/gofakeit"
	"log"
	"strings"
	"syscall"
	"time"
)

func main() {
	stream := flag.String("stream", "", "The name of your stream")
	flag.Parse()

	if strings.EqualFold(*stream, "") {
		flag.Usage()
		syscall.Exit(1)
	}

	newSession, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	client := kinesis.New(newSession)

	for {
		records := make([]*kinesis.PutRecordsRequestEntry, 0)

		for i := 0; i < 200; i++ {
			randomLog := fmt.Sprintf(`%s - - [%s] "%s HTTP/1.0" %d %d "%s" "%s"`,
				gofakeit.IPv4Address(),
				gofakeit.DateRange(time.Now().Add(-1*1*time.Hour), time.Now()).Format("2/Jan/2006:15:04:05 -0700"),
				gofakeit.HTTPMethod(),
				gofakeit.StatusCode(),
				gofakeit.Uint16(),
				gofakeit.URL(),
				gofakeit.UserAgent())

			partitionKey := gofakeit.UUID()
			records = append(records, &kinesis.PutRecordsRequestEntry{
				Data:         []byte(randomLog),
				PartitionKey: &partitionKey,
			})

			fmt.Println(randomLog)
		}

		if _, err := client.PutRecords(&kinesis.PutRecordsInput{
			Records:    records,
			StreamName: stream,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
