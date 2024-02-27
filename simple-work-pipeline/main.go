/**
Overall flow is:
1. Do the setup
2. Read file
3. Process
4. Output

Usage: go run main.go <filepath> [range of services names]
**/

package main

import (
	"context"
	"log"
	"os"
)

type Processor struct {
	ctx      context.Context
	services [](func(text string) string)
}

func (pr *Processor) RegisterServices(services []string) {
	for _, service := range services {
		switch service {
		case "upper":
			pr.services = append(pr.services, NewUpperService())
		case "reverse":
			pr.services = append(pr.services, NewReverseService())
		case "remove-white-space":
			pr.services = append(pr.services, NewRemoveWhiteSpaceService())
		default:
			log.Fatalf("service %s not supported\n", service)
		}
	}
}

func (pr *Processor) Apply(text string) string {

	for _, service := range pr.services {
		text = service(text)
	}

	return text
}

func NewProcessor() *Processor {
	pr := &Processor{ctx: context.Background()}
	return pr
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("usage: go run reader.go <input_file_path> <output_file_path> [range of services names]")
		return
	}

	inputFilePath := os.Args[1]
	reader := SelectReader(inputFilePath)
	if reader == nil {
		return
	}

	reader.Read(inputFilePath)

	outputFilePath := os.Args[2]
	writer := SelectWriter(outputFilePath)
	if writer == nil {
		return
	}

	// Create a processor
	processor := NewProcessor()

	var services []string
	if len(os.Args) > 3 {
		services = append(services, os.Args[3:]...)
	}

	// Register services
	processor.RegisterServices(services)

	var (
		res string
		err error
	)
	for {
		text, exist := reader.Next()
		if !exist {
			break
		}

		res = processor.Apply(text)
		err = writer.Write(res)
		if err != nil {
			log.Fatalf("write to file failed, err: %s\n", err)
			return
		}
	}

	writer.Close()
	reader.Close()
}
