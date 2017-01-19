func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	myItems := []string{"item1", "item2", "item3"}
	a, _ := json.Marshal(myItems)

	w.Write(a)
	return
}
