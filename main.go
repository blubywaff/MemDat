package main

import (
	"fmt"
	"time"
)

type Comment struct {
	CommentStem
	Replies []string `json:"Replies" bson:"Replies" memdat:"Replies"`
}

type CommentStem struct {
	ID   string `json:"ID" bson:"ID"`
	Body string `json:"Body" bson:"Body"`
}

type FullComment struct {
	CommentStem
	Replies []FullComment
}

func main() {
	/*
		fmt.Println((&Comment{CommentStem{"ee", "no u"}, []string{"hello", "I agree with you"}}).MemDatConvert())
		fmt.Println("e\n\r\te")
		fmt.Println(strings.Contains("e\n\te", "\r"))

		var data interface{}
		str := "{\n\"EE\": [\n\t\"\\\"\\\"\\\"\\n\\\"\\\"\"\n\t]\n}"
		fmt.Println(str)
		fmt.Println(json.Unmarshal([]byte(str), &data))
		fmt.Println(data)
		data = nil
		fmt.Println("{\"EE\": [\"\\\"\\\\\\\"\"]}")
		fmt.Println(json.Unmarshal([]byte("{\"EE\": [\"\\\\\"]}"), &data))
		fmt.Println(data)
	*/

	start := time.Now().UnixNano()

	database := *NewDatabase()
	for i := 0; i < 5; i++ {
		var cm *Comment
		cm = &Comment{CommentStem{"xID" + database.generateId(), "xBody"}, []string{database.generateId(), "xx"}}
		fmt.Println("ID--", cm.Replies[0])
		database.Add(*cm)
	}
	//database.Add([]int{1, 2, 3})
	fmt.Println(database.Documents[0])
	fmt.Println(database.Indexes[0].findDocument(database.Documents[0]["ObjectId"].(string)))
	(*database.Indexes[0].findDocument(database.Documents[0]["ObjectId"].(string)))["ObjectId"] = "0"
	fmt.Println(database.Indexes[0].findDocument(database.Documents[0]["ObjectId"].(string)))
	fmt.Println(database.Indexes[0].Index)
	for i, _ := range database.findIndex("ObjectId").Index {
		fmt.Println(database.Indexes[0].Index[i].Value, (*database.Indexes[0].Index[i].Document)["Replies"])
	}

	fmt.Println(time.Now().UnixNano() - start)

	//database.Add([]string{"Hello", "21"})

	//fmt.Println(func(i interface{}) interface{} { return i }(Comment{CommentStem{"ee", "no u"}, []string{"hello", "I agree with you"}}))
	/*
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprint(w, "EE")
			fmt.Println(req.Header)
		})
		log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) { mux.ServeHTTP(w, req) })))
	*/
}
