package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/brianvoe/gofakeit"
)

func main() {
	userstream := flag.String("stream", "", "The name of your stream")
	userstreams := flag.String("streams", "", "A comma separated list of the names of your streams")
	var streams []*string
	flag.Parse()

	if strings.EqualFold(*userstream, "") && strings.EqualFold(*userstreams, "") {
		flag.Usage()
		syscall.Exit(1)
	}

	if !strings.EqualFold(*userstream, "") {
		streams = append(streams, userstream)
		syscall.Exit(1)
	} else {
		splitstreams := strings.Split(*userstreams, ",")
		for _, s := range splitstreams {
			clean := strings.TrimSpace(s)
			streams = append(streams, &clean)
		}
	}

	newSession, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	client := kinesis.New(newSession)

	for {
		for _, stream := range streams {
			records := make([]*kinesis.PutRecordsRequestEntry, 0)
			fmt.Println(*stream)

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
}
