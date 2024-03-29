package payment

import (
	"testing"

	"github.com/BrunoDM2943/pock_bank-api/domain"
)

func BenchmarkProcessPayments(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processPayments(mockList(10000))
	}
}

func BenchmarkProcessPaymentsFaster(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processPaymentsFaster(mockListPointer(10000))
	}
}

func mockList(qtd int) []domain.Payment {
	payments := make([]domain.Payment, 0)
	for i := 0; i < qtd; i++ {
		p := domain.Payment{
			ID: i,
		}
		payments = append(payments, p)
	}
	return payments
}

func mockListPointer(qtd int) []*domain.Payment {
	payments := make([]*domain.Payment, 0)
	for i := 0; i < qtd; i++ {
		p := &domain.Payment{
			ID: i,
		}
		payments = append(payments, p)
	}
	return payments
}
