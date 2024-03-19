package services

import (
	"github.com/PiotrFerenc/mash2/api/types"
	"github.com/PiotrFerenc/mash2/internal/queues"
	"github.com/goccy/go-json"
	"log"
)

type PipelineService interface {
	Run(pipeline *types.Pipeline) error
}

type pipelineService struct {
	queue          queues.MessageQueue
	processService ProcessService
}

func CreatePipelineService(queue queues.MessageQueue, processService ProcessService) PipelineService {
	s := func(message queues.Message) {

		index := message.CurrentStage.Order
		if index < len(message.Pipeline.Stages) {
			current := message.Pipeline.Stages[index]
			err := queue.AddStageToQueue(queues.Message{
				CurrentStage: current,
				Pipeline:     message.Pipeline,
			})
			if err != nil {
				panic(err)
			}
		} else {
			log.Printf("Done")
		}

	}

	f := func(message queues.Message) {
		log.Printf(" fail [x] %s", message.CurrentStage.Name)
	}

	go func(onSucces func(queues.Message), onFail func(queues.Message)) {
		ss, _ := queue.WaitingForSucceedStage()
		fs, _ := queue.WaitingForFailedStage()

		var forever chan struct{}

		go func() {
			for d := range ss {
				var message queues.Message
				err := json.Unmarshal(d.Body, &message)
				if err != nil {
					panic(err)
				}

				onSucces(message)
			}
		}()

		go func() {
			for d := range fs {
				var message queues.Message
				err := json.Unmarshal(d.Body, &message)
				if err != nil {
					panic(err)
				}

				onFail(message)
			}
		}()

		<-forever

	}(s, f)

	return &pipelineService{
		queue:          queue,
		processService: processService,
	}
}

func (p *pipelineService) Run(pipeline *types.Pipeline) error {

	err := p.queue.AddStageToQueue(queues.Message{
		CurrentStage: pipeline.Stages[0],
		Pipeline:     *pipeline,
	})
	if err != nil {
		return err
	}

	return nil
}
