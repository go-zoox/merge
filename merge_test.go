package merge

import "testing"

func TestMergeStructsTwo(t *testing.T) {
	type Data struct {
		Name string
		Age  int
	}

	target := Data{Name: "John", Age: 30}
	source1 := Data{Name: "Jane", Age: 25}

	if err := Merge(&target, &source1); err != nil {
		t.Fatal(err)
	}

	if target.Name != "Jane" {
		t.Errorf("expected %s, got %s", "Jane", target.Name)
	}

	if target.Age != 25 {
		t.Errorf("expected %d, got %d", 25, target.Age)
	}
}

func TestMergeStructsThree(t *testing.T) {
	type Data struct {
		Name string
		Age  int
	}

	target := Data{Name: "John", Age: 30}
	source1 := Data{Name: "Jane", Age: 25}
	source2 := Data{Name: "Joe", Age: 20}

	if err := Merge(&target, &source1, &source2); err != nil {
		t.Fatal(err)
	}

	if target.Name != "Joe" {
		t.Errorf("expected %s, got %s", "Joe", target.Name)
	}

	if target.Age != 20 {
		t.Errorf("expected %d, got %d", 20, target.Age)
	}
}

func TestMergeStructsEmpty(t *testing.T) {
	type Data struct {
		Name string
		Age  int
	}

	target := Data{Name: "John", Age: 30}
	source1 := Data{Name: "Jane", Age: 25}
	source2 := Data{Name: "Joe"}

	if err := Merge(&target, &source1, &source2); err != nil {
		t.Fatal(err)
	}

	if target.Name != "Joe" {
		t.Errorf("expected %s, got %s", "Joe", target.Name)
	}

	if target.Age != 0 {
		t.Errorf("expected %d, got %d", 0, target.Age)
	}
}

func TestMergeStructsNil(t *testing.T) {
	type Data struct {
		Name string
		Age  int
	}

	target := Data{Name: "John", Age: 30}
	source1 := Data{Name: "Jane", Age: 25}

	if err := Merge(&target, &source1, nil); err != nil {
		t.Fatal(err)
	}

	if target.Name != "Jane" {
		t.Errorf("expected %s, got %s", "Jane", target.Name)
	}

	if target.Age != 25 {
		t.Errorf("expected %d, got %d", 25, target.Age)
	}
}

func TestMergeStructsNil2(t *testing.T) {
	type Data struct {
		Name string
		Age  int
	}

	target := Data{Name: "John", Age: 30}
	source1 := Data{Name: "Jane", Age: 25}

	if err := Merge(&target, nil, &source1); err != nil {
		t.Fatal(err)
	}

	if target.Name != "Jane" {
		t.Errorf("expected %s, got %s", "Jane", target.Name)
	}

	if target.Age != 25 {
		t.Errorf("expected %d, got %d", 25, target.Age)
	}
}

func TestMergeStructsNest(t *testing.T) {
	type Address struct {
		Street string
		City   string
	}

	type Data struct {
		Name    string
		Age     int
		Address Address
	}

	target := Data{Name: "John", Age: 30, Address: Address{Street: "123 Main St", City: "Anytown"}}
	source1 := Data{Name: "Jane", Age: 25, Address: Address{Street: "456 Main St", City: "Anytown2"}}

	if err := Merge(&target, nil, &source1); err != nil {
		t.Fatal(err)
	}

	if target.Name != "Jane" {
		t.Errorf("expected %s, got %s", "Jane", target.Name)
	}

	if target.Age != 25 {
		t.Errorf("expected %d, got %d", 25, target.Age)
	}

	if target.Address.Street != "456 Main St" {
		t.Errorf("expected %s, got %s", "456 Main St", target.Address.Street)
	}

	if target.Address.City != "Anytown2" {
		t.Errorf("expected %s, got %s", "Anytown2", target.Address.City)
	}
}
