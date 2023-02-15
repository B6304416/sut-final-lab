package entity

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeEmployee(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("test positive", func(t *testing.T) {
		emID := "M123456789"
		em := Employee{
			Name:       "",
			Email:      "pon@gmail.com",
			EmployeeID: emID,
		}
		ok, err := govalidator.ValidateStruct(em)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("EmployeeID: %s does not validate as matches(^[JMS]\\d{8}$)", emID)))
	})

}
