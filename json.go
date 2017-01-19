
  //S1 OMIT
        import "encoding/json"
       
	bolB, _ := json.Marshal(true)
	intB, _ := json.Marshal(1)
	fltB, _ := json.Marshal(2.34)
	strB, _ := json.Marshal("gopher")
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
//S2 OMIT
        type Response1 struct {
         	Page   int
         	Fruits []string
        }

        res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}
	}

        res1B, _ := json.Marshal(res1D)
    //S3 OMIT
        type Response2 struct {
        	Page   int      `json:"page"` //tag
	        Fruits []string `json:"fruits"`
        }

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res2B, _ := json.Marshal(res2D)

     //S4 OMIT
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
        json.Unmarshal([]byte(str), &res)

        enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
//S5 OMIT
