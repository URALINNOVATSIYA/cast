package cast

import (
	"testing"
)

type person struct {
	Name  string
	Age   int
	Sex   bool
	Data  any
	Score [][]int
	foo   string
}

type human struct {
	Name  string
	Age   float64
	Sex   string
	Data  any
	Score [][]float32
	Foo   bool
	boo   uint
}

type personContact struct {
	Phone string
	Email string
	Url   string
}

type complexPerson struct {
	Name     string
	Contacts personContact
	Family   []complexPerson
}

type humanContact struct {
	Phone uint
	Email string
}

type complexHuman struct {
	Name     string
	Contacts humanContact
	Family   []complexHuman
}

func TestAsStruct(t *testing.T) {
	t.Run("cast to simple struct", func(t *testing.T) {
		tests := []castTest[person]{
			{nil, person{}, ""},
			{
				map[string]any{
					"name":  "Vasya Pupkin",
					"Age":   123,
					"sex":   true,
					"Data":  "test",
					"score": [][]string{{"1", "2"}, {"3"}},
					"Foo":   1.0,
				},
				person{
					Name:  "Vasya Pupkin",
					Age:   123,
					Sex:   true,
					Data:  "test",
					Score: [][]int{{1, 2}, {3}},
					foo:   "",
				},
				"",
			},
			{
				human{
					Name:  "Vasya Pupkin",
					Age:   30.5,
					Sex:   "true",
					Data:  []any{1, 2, 3},
					Score: [][]float32{{1.5, 2.4}, {3.7}},
					Foo:   true,
					boo:   7,
				},
				person{
					Name:  "Vasya Pupkin",
					Age:   30,
					Sex:   true,
					Data:  []any{1, 2, 3},
					Score: [][]int{{1, 2}, {3}},
					foo:   "",
				},
				"",
			},
			{
				[]int{1, 2},
				person{},
				"failed to cast []int to cast.person",
			},
		}
		runCastTests(t, "AsStruct[person]", AsStruct[person], tests)
	})

	t.Run("cast to complex struct", func(t *testing.T) {
		tests := []castTest[complexPerson]{
			{nil, complexPerson{}, ""},
			{
				map[string]any{
					"name": "Vasya Pupkin",
					"contacts": map[string]string{
						"phone": "+7555112233",
						"email": "foo@gmail.com",
						"url":   "https://google.com",
					},
					"family": []map[string]any{
						{
							"name": "Goga",
							"contacts": map[string]string{
								"email": "goga@gmail.com",
							},
						},
						{
							"name": "Mira",
							"contacts": map[string]string{
								"url": "https://foo.boo",
							},
							"family": []map[string]any{
								{
									"name": "Pet",
								},
							},
						},
					},
				},
				complexPerson{
					Name: "Vasya Pupkin",
					Contacts: personContact{
						Phone: "+7555112233",
						Email: "foo@gmail.com",
						Url:   "https://google.com",
					},
					Family: []complexPerson{
						{
							Name: "Goga",
							Contacts: personContact{
								Email: "goga@gmail.com",
							},
						},
						{
							Name: "Mira",
							Contacts: personContact{
								Url: "https://foo.boo",
							},
							Family: []complexPerson{
								{
									Name: "Pet",
								},
							},
						},
					},
				},
				"",
			},
			{
				complexHuman{
					Name: "Vasya Pupkin",
					Contacts: humanContact{
						Phone: 1234567890,
						Email: "foo@gmail.com",
					},
					Family: []complexHuman{
						{
							Name: "Goga",
							Contacts: humanContact{
								Email: "goga@gmail.com",
							},
						},
						{
							Name: "Mira",
							Contacts: humanContact{
								Phone: 555666777,
							},
							Family: []complexHuman{
								{
									Name: "Pet",
								},
							},
						},
					},
				},
				complexPerson{
					Name: "Vasya Pupkin",
					Contacts: personContact{
						Phone: "1234567890",
						Email: "foo@gmail.com",
					},
					Family: []complexPerson{
						{
							Name: "Goga",
							Contacts: personContact{
								Phone: "0",
								Email: "goga@gmail.com",
							},
						},
						{
							Name: "Mira",
							Contacts: personContact{
								Phone: "555666777",
							},
							Family: []complexPerson{
								{
									Name: "Pet",
									Contacts: personContact{
										Phone: "0",
									},
								},
							},
						},
					},
				},
				"",
			},
			{
				map[string]any{
					"name":   "Vasya Pupkin",
					"Family": []any{"foo"},
				},
				complexPerson{
					Name: "Vasya Pupkin",
				},
				"failed to cast string to cast.complexPerson",
			},
		}
		runCastTests(t, "AsStruct[complexPerson]", AsStruct[complexPerson], tests)
	})
}
