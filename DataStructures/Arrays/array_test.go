// This is a file to test the Array struct to run this you can use go test

// testing package help in automating it

// if the argument t *testing.T is given then go will run the function automatically in go test

// Tips: To get all the log values of a test you can run go test -v , then it will give log also

package array

import "testing"

func TestAllBasicMethodOnArray(t *testing.T) {
	cap := 5
	arr := NewArray(uint(cap))

	arr.Print()

	for i := 0; i < cap-2; i++ {
		err := arr.Insert(uint(i), i+1)

		if err != nil {
			t.Fatal(err.Error())
		}
	}

	if true {
		err := arr.Insert(uint(3), 100)
		arr.Print()
		if err != nil {
			t.Errorf("err:%+v", err)
		}
	}

	if true {
		err := arr.InsertTailend(999)
		arr.Print()
		if err != nil {
			t.Errorf("err:%+v", err)
		}
	}

	arr.Print()
	t.Log(arr.Len())
	t.Log(arr.isIndexOutofRange(12))

}

func TestFind(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))

	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}

	arr.Print()

	t.Log(arr.Find(8))
	t.Log(arr.Find(0))
	t.Log(arr.Find(12))

}
