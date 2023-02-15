package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("test positive", func(t *testing.T) {
		em := Employee{
			Name:       "natthapon",
			Email:      "pon@gmail.com",
			EmployeeID: "M12345678",
		}
		ok, err := govalidator.ValidateStruct(em)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("name not blank"))
	})

}
