package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Nics",
		Price: 1.00,
		SKU:   "abs-abs-abs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
