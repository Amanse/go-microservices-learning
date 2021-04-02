package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "nics",
		Price: 1.00,
		SKU:   "supa-gay-lad",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
