package event

type EventServer struct {
	// ConsumerProductionController *controller.ConsumerProductionController
	// ConsumerService              event.IConsumerService
}

// func NewEventServer(ConsumerService event.ConsumerService,
// 	consumerProductionController controller.ConsumerProductionController) *EventServer {
// 	slog.InfoContext(context.Background(), "Creating Event Server...")

// cpc := controller.NewConsumerProductionController(container)

// return &EventServer{
// ConsumerProductionController: cpc,
// }
// }

// func (eventServer *EventServer) Consume(ctx context.Context) error {
// 	slog.InfoContext(ctx, "Starting Event Server...")

// 	// err2 := eventServer.ConsumerProductionController.Execute(ctx)
// 	// if err2 != nil {
// 	// 	slog.ErrorContext(ctx, "Error starting Event Server: ", err2)
// 	// 	return err2
// 	// }

// 	slog.InfoContext(ctx, "Event Server started successfully")

// 	return nil
// }
