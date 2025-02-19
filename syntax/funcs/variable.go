package main

func YourName(name string, aliases ...string) {
	if len(aliases) > 0 {
		println(name, aliases[len(aliases)-1])
	}
}
func YourNameInvoke() {
	YourName("Lee")
	YourName("Lee", "A")
	YourName("Lee", "A", "B")
	YourName("Lee", "A", "B", "C", "D")
}
