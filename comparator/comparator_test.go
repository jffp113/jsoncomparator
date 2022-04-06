package comparator

import (
	"bytes"
	"testing"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestCompareJSON(t *testing.T) {

	t.Log("Given the need to compare two arrays with json objects a user with admin method.")
	{

		testID := 0
		t.Logf("\tTest %d:\tWhen a user tries to compare two identical arrays with json objects", testID)
		{

			file1 := `[
				{
					"id": "123",
					"name": "alice"	
				},
				{
					"id": "124",
					"name": "bob"	
				}
			]`

			file2 := `[
			{
				"name": "alice",	
				"id": "123"
			},
			{
				"name": "bob",
				"id": "124"	
			}
		]`

			r1 := bytes.NewReader([]byte(file1))
			r2 := bytes.NewReader([]byte(file2))

			equals, err := CompareJSON(r1, r2)

			if err != nil {
				t.Fatalf("\t%s\tTest %d:\t Should not return a error.", Failed, testID)
			}
			t.Logf("\t%s\tTest %d:\t Should not return a error.", Success, testID)

			if !equals {
				t.Fatalf("\t%s\tTest %d:\t Should be equal.", Failed, testID)
			}
			t.Logf("\t%s\tTest %d:\t Should be equal.", Success, testID)
		}

		testID = testID + 1
		t.Logf("\tTest %d:\tWhen a user tries to compare two arrays with json objects with different field values", testID)
		{

			file1 := `[
				{
					"id": "130",
					"name": "alice"	
				},
				{
					"id": "124",
					"name": "bob"	
				}
			]`

			file2 := `[
			{
				"name": "alice",	
				"id": "123"
			},
			{
				"name": "bob",
				"id": "124"	
			}
		]`

			r1 := bytes.NewReader([]byte(file1))
			r2 := bytes.NewReader([]byte(file2))

			equals, err := CompareJSON(r1, r2)

			if err != nil {
				t.Fatalf("\t%s\tTest %d:\t Should not return a error.", Failed, testID)
			}
			t.Logf("\t%s\tTest %d:\t Should not return a error.", Success, testID)

			if equals {
				t.Fatalf("\t%s\tTest %d:\t Should not be equal.", Failed, testID)
			}
			t.Logf("\t%s\tTest %d:\t Should not be equal.", Success, testID)
		}
		testID = testID + 1
		t.Logf("\tTest %d:\tWhen a user tries to compare two arrays with json objects with different number of fields", testID)
		{

			file1 := `[
				{
					"id": "123",
					"name": "alice"	
				},
				{
					"id": "124",
					"name": "bob"	
				}
			]`

			file2 := `[
				{
					"name": "alice",	
					"id": "123"
				},
				{
					"name": "bob",
					"id": "124",	
					"phone_number": "960000000"
				}
			]`

			r1 := bytes.NewReader([]byte(file1))
			r2 := bytes.NewReader([]byte(file2))

			equals, err := CompareJSON(r1, r2)

			if err != nil {
				t.Fatalf("\t%s\tTest %d:\t Should not return a error.", Failed, testID)
			}
			t.Logf("\t%s\tTest %d:\t Should not return a error.", Success, testID)

			if equals {
				t.Fatalf("\t%s\tTest %d:\t Should not be equal.", Failed, testID)
			}
			t.Logf("\t%s\tTest %d:\t Should not be equal.", Success, testID)
		}
		testID = testID + 1
		t.Logf("\tTest %d:\tWhen a user tries to compare two arrays with json objects not correctly built", testID)
		{

			file1 := `[
				{
					"id": "123",
					"name": "alice"	
				} missing comma here :)
				{
					"id": "124",
					"name": "bob"	
				}
			]`

			file2 := `[
				{
					"name": "alice",	
					"id": "123"
				} missing comma here :)
				{
					"name": "bob",
					"id": "124"	
				}
			]`

			r1 := bytes.NewReader([]byte(file1))
			r2 := bytes.NewReader([]byte(file2))

			_, err := CompareJSON(r1, r2)

			if err == nil {
				t.Fatalf("\t%s\tTest %d:\t Should return a error.", Failed, testID)
			}
			t.Logf("\t%s\tTest %d:\t Should return a error.", Success, testID)

		}
	}

}
