package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/loggingsearch"
)

var registros int
var totalLinhas int

// Função que grava o arquivo de log
func writeFile(filename string, logs []loggingsearch.SearchResult) {

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, value := range logs {
		jsonBytes, err := json.Marshal(value)
		if err != nil {
			log.Fatalf("Failed to marshal entries: %v", err)
		}
		_, err = file.Write(jsonBytes)
		_, err = file.WriteString("\n")
	}
}

var inputDatainicial string
var inputDatafinal string
var compartment string

// Função que obter os logs
func getLog(request string) {

	if len(os.Args) < 2 {
		log.Fatal("Você precisa informar a data inicial do log. Exemplo: 2023-01-31T00:00:00")
	}

	if len(os.Args) < 3 {
		log.Fatal("Você precisa informar a data do log. Exemplo: 2023-01-31T00:00:00")
	}

	if len(os.Args) < 3 {
		log.Fatal("Você precisa informar o OCID do compartimento para buscar os logs.")
	}

	dataInicial, err := time.Parse(time.RFC3339, inputDatainicial+".600Z")
	if err != nil {
		log.Fatal("Formato de data inválido. A data deve estar no formato 2023-01-31T00:00:00")
		return
	}

	dataFinal, err := time.Parse(time.RFC3339, inputDatafinal+".600Z")
	if err != nil {
		log.Fatal("Formato de data inválido. A data deve estar no formato 2023-01-31T00:00:00")
		return
	}

	client, err := loggingsearch.NewLogSearchClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	start := dataInicial
	end := dataFinal

	var req loggingsearch.SearchLogsRequest

	if request != "" {
		req = loggingsearch.SearchLogsRequest{OpcRequestId: common.String("RIJJ87OXLQBMPBRN6ZWO/OpcRequestIdExample/UniqueRequestId"),
			SearchLogsDetails: loggingsearch.SearchLogsDetails{SearchQuery: common.String(`search "` + compartment + `"`),
				TimeStart:         &common.SDKTime{Time: start},
				TimeEnd:           &common.SDKTime{Time: end},
				IsReturnFieldInfo: common.Bool(false)},
			Page:  &request,
			Limit: common.Int(1000)}
	} else {
		req = loggingsearch.SearchLogsRequest{OpcRequestId: common.String("RIJJ87OXLQBMPBRN6ZWO/OpcRequestIdExample/UniqueRequestId"),
			SearchLogsDetails: loggingsearch.SearchLogsDetails{SearchQuery: common.String(`search "` + compartment + `"`),
				TimeStart:         &common.SDKTime{Time: start},
				TimeEnd:           &common.SDKTime{Time: end},
				IsReturnFieldInfo: common.Bool(false)},
			Limit: common.Int(1000)}
	}
	resp, err := client.SearchLogs(context.Background(), req)
	helpers.FatalIfError(err)
	writeFile("log-"+strings.Replace(inputDatainicial, ":", "-", -1)+"_"+strings.Replace(inputDatafinal, ":", "-", -1)+".jsonl", resp.SearchResponse.Results)
	if resp.OpcNextPage == nil {
		fmt.Println("Concluída a geração de log")
	} else {
		registros = registros + 1000
		if registros >= 50000 {
			totalLinhas = totalLinhas + registros
			registros = 0
			now := time.Now()
			fmt.Println(now.Format("02/01/2006 15:04:05"), "- Processadas", totalLinhas, "Linhas.")
		}
		getLog(*resp.OpcNextPage)
	}
}

func main() {
	flag.StringVar(&inputDatainicial, "data-inicial", "", "A data inicial no formato 2023-01-31T00:00:00")
	flag.StringVar(&inputDatafinal, "data-final", "", "A data final no formato 2023-01-31T00:00:00")
	flag.StringVar(&compartment, "compartment", "", "O ocid do compartimento")
	flag.Parse()
	getLog("")
}
