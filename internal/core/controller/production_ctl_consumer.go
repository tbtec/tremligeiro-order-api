package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type ConsumerProductionController struct {
	usc *usecase.UscUpdateOrder
	// ConsumerService event.IConsumerService
}

func NewConsumerProductionController(container *container.Container) *ConsumerProductionController {
	return &ConsumerProductionController{
		usc: usecase.NewUscUpdateOrder(
			gateway.NewOrderGateway(container.OrderRepository, container.ProducerService),
		),
		// ConsumerService: &container.ConsumerService,
	}
}

func (ctl *ConsumerProductionController) Execute(ctx context.Context, orderId string, newStatus string) error {

	// order, _ := ctl.ConsumerService.ConsumeMessage(ctx)

	// return ctl.usc.Update(ctx, order.ID, order.Status)
	return nil
}
