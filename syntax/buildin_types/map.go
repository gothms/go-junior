package main

func Map() {
	m1 := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
		"k5": "v5",
	}
	println(m1)

	m2 := make(map[string]string, 4)
	m2["k3"] = "v3"
	m2["k4"] = "v4"

	val, ok := m1["k1"]
	if ok {
		println(val)
	}

	delete(m1, "k3")
	for k, v := range m1 {
		println(k, v)
	}

	clear(m2)
	for k, v := range m2 {
		println(k, v)
	}
}
