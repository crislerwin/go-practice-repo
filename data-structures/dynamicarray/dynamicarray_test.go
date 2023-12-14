package dynamicarray

import (
	"reflect"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	numbers := DynamicArray{}

	t.Run("Check empty Dynamic Array", func(t *testing.T) {
		if numbers.IsEmpty() != true {
			t.Errorf("Expected true but got %v", numbers.IsEmpty())
		}
	})

	numbers.Add(10)
	numbers.Add(20)
	numbers.Add(30)
	numbers.Add(40)
	numbers.Add(50)

	t.Run("Add Element into Dynamic Array", func(t *testing.T) {
		if numbers.IsEmpty() != false {
			t.Errorf("Expected be false but got %v", numbers.IsEmpty())
		}
		var res []any = []any{10, 20, 30, 40, 50}
		if !reflect.DeepEqual(numbers.GetData(), res) {
			t.Errorf("Expected be [10, 20, 30, 40, 50] but got %v", numbers.GetData())
		}
	})
	t.Run("Remove in Dynamic Array", func(t *testing.T) {
		if numbers.IsEmpty() != false {
			var res []any
			res = append(res, 10)
			res = append(res, 20)
			res = append(res, 30)
			res = append(res, 40)
			res = append(res, 50)

			err := numbers.Remove(1)
			if err != nil {
				t.Errorf("Expected be [10, 30, 40, 50] but got an Error %v", err)
			}
			if !reflect.DeepEqual(numbers.GetData(), res) {
				t.Errorf("Expected be [10, 30, 40, 50], but got %v", numbers.GetData())
			}
		}
	})
	t.Run("Get in Dynamic Array", func(t *testing.T) {
		if numbers.IsEmpty() != false {
			t.Errorf("Expected be false but got %v", numbers.IsEmpty())
		}
		getOne, _ := numbers.Get(2)
		if getOne != 40 {
			t.Errorf("Expected be 40 but got %v", getOne)
		}
	})
	t.Run("Put on Dynamic Array", func(t *testing.T) {
		if numbers.IsEmpty() != false {
			t.Errorf("Expected be false but got %v", numbers.IsEmpty())
		}
		err := numbers.Put(0, 100)
		if err != nil {
			t.Errorf("Expected be [100, 30, 40, 50], but got an error %v", err)
		}

		getOne, _ := numbers.Get(0)
		if getOne != 100 {
			t.Errorf("Expected be 100 but got %v", getOne)
		}
	})
}
