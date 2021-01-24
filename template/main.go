package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Info struct {
	Gender string
	Name string
	Age int
}

func Greet(name string) (string, error) {
	return "How are you " + name, nil
}
func Ask() string {
	return "Do you have launch ?"
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// transport parameters into template
	t, err := template.ParseFiles("templates/halo.tmpl")
	if err != nil {
		fmt.Println("template load fail .", err)
		return
	}

	var per1 = Info{"male", "per1", 3}
	var per2 = map[string]interface{}{
		"gender": "male",
		"name": "per2",
		"age": 5,
	}
	var coup = map[string]interface{} {
		"person1": per1,
		"person2": per2,
		"numbers": []int{1, 3, 5, 7, 9},
	}

	err = t.Execute(w, coup)
	if err != nil {
		fmt.Println("template execute fail .", err)
	}
}

func describe(w http.ResponseWriter, r *http.Request) {
	// transport function into template
	t := template.New("introduce.tmpl")
	t.Funcs(template.FuncMap{
		"fun1": Greet,
		"fun2": Ask,
	})
	t, err := t.ParseFiles("templates/introduce.tmpl")
	if err != nil {
		fmt.Println("template load fail .", err)
		return
	}

	t.Execute(w, " baby dog !")
}

func nest(w http.ResponseWriter, r *http.Request) {
	// test template nest
	t, err := template.ParseFiles("templates/nest.tmpl","templates/base1.tmpl")
	if err != nil {
		fmt.Println("template load fail .", err)
		return
	}
	t.Execute(w, " dog baby !")
}

func inherit1(w http.ResponseWriter, r *http.Request) {
	// test template inherit
	t, err := template.ParseFiles("templates/baseT.tmpl", "templates/inherit1.tmpl")
	if err != nil {
		fmt.Println("template load fail .", err)
		return
	}
	err = t.ExecuteTemplate(w, "inherit1.tmpl", "inherit 1")
	if err != nil {
		fmt.Println("template render fail .", err)
		return
	}
}
func inherit2(w http.ResponseWriter, r *http.Request) {
	// test template inherit
	t, err := template.ParseFiles("templates/baseT.tmpl", "templates/inherit2.tmpl")
	if err != nil {
		fmt.Println("template load fail .", err)
		return
	}
	err = t.ExecuteTemplate(w, "inherit2.tmpl", "inherit 2")
	if err != nil {
		fmt.Println("template render fail .", err)
		return
	}
}
func inherit3(w http.ResponseWriter, r *http.Request) {
	// test template inherit
	t, err := template.ParseFiles("templates/baseT.tmpl", "templates/inherit3.tmpl")
	if err != nil {
		fmt.Println("template load fail .", err)
		return
	}
	err = t.ExecuteTemplate(w, "inherit3.tmpl", "inherit 3")
	if err != nil {
		fmt.Println("template render fail .", err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	
	t, err := template.New("index.tmpl").
						Delims("{[", "]}").
						ParseFiles("templates/index.tmpl")
	if err != nil {
		fmt.Println("template load fail .", err)
		return
	}
	err = t.Execute(w, "INDEX")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/introduce", describe)
	http.HandleFunc("/nest", nest)
	http.HandleFunc("/inherit1", inherit1)
	http.HandleFunc("/inherit2", inherit2)
	http.HandleFunc("/inherit3", inherit3)
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("failed")
		return
	}
}