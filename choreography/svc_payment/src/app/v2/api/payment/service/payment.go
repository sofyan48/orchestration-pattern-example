package service

import (
	"strconv"
	"time"

	"github.com/sofyan48/svc_payment/src/app/v2/api/payment/entity"
	"github.com/sofyan48/svc_payment/src/app/v2/api/payment/event"
	"github.com/sofyan48/svc_payment/src/utils/kafka"
)

// PaymentService ...
type PaymentService struct {
	Event event.PaymentEventInterface
	Kafka kafka.KafkaLibraryInterface
	ready chan bool
}

// PaymentServiceHandler ...
func PaymentServiceHandler() *PaymentService {
	return &PaymentService{
		Event: event.PaymentEventHandler(),
		Kafka: kafka.KafkaLibraryHandler(),
		ready: make(chan bool),
	}
}

// PaymentServiceInterface ...
type PaymentServiceInterface interface {
	PaymentCreateService(payload *entity.PaymentRequest) (*entity.PaymentResponses, error)
	PaymentUpdateOrder(OrderUUID string, payload *entity.PaymentPaidRequest) (*entity.PaymentResponses, error)
	PaymentGetStatus(uuid string) (interface{}, error)
	ListPayment(limit, page int) (interface{}, error)
}

// PaymentCreateService ...
func (service *PaymentService) PaymentCreateService(payload *entity.PaymentRequest) (*entity.PaymentResponses, error) {
	now := time.Now()
	eventPayload := &entity.PaymentEvent{}
	eventPayload.Action = "payment_save"
	eventPayload.CreatedAt = &now
	data := map[string]interface{}{
		"bank_account_number": payload.BankAccountNumber,
		"id_payment_model":    payload.IDPaymentModel,
		"id_payment_status":   payload.IDPaymentStatus,
		"inquiry_number":      payload.InquiryNumber,
		"nm_bank":             payload.NMBank,
		"payment_total":       payload.PaymentTotal,
	}
	eventPayload.Data = data
	event, err := service.Event.PaymentCreateEvent(eventPayload)
	if err != nil {
		return nil, err
	}
	result := &entity.PaymentResponses{}
	result.UUID = event.UUID
	result.Event = event
	result.CreatedAt = event.CreatedAt
	return result, nil
}

// PaymentUpdateOrder ...
func (service *PaymentService) PaymentUpdateOrder(OrderUUID string, payload *entity.PaymentPaidRequest) (*entity.PaymentResponses, error) {
	now := time.Now()
	eventPayload := &entity.PaymentEvent{}
	eventPayload.Action = "payment_order"
	eventPayload.CreatedAt = &now
	payTotal := strconv.Itoa(payload.PaymentTotal)
	data := map[string]interface{}{
		"uuid_order":          OrderUUID,
		"payment_total":       payTotal,
		"payment_status":      payload.PaymentStatus,
		"bank_account_number": payload.BankAccountNumber,
	}
	eventPayload.Data = data
	event, err := service.Event.PaymentCreateEvent(eventPayload)
	if err != nil {
		return nil, err
	}
	result := &entity.PaymentResponses{}
	result.UUID = event.UUID
	result.Event = event
	result.CreatedAt = event.CreatedAt
	return result, nil
}

// PaymentGetStatus ...
func (service *PaymentService) PaymentGetStatus(uuid string) (interface{}, error) {
	return nil, nil
}

// ListPayment ...
func (service *PaymentService) ListPayment(limit, page int) (interface{}, error) {
	return nil, nil
}
