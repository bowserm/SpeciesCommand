package main

import(
		"net/http"
		"log"
		"os"
	  )

const resp = `<html>
<head>
<title>Species Command</title>
</head>
<body>
<h1>Species</h1>
<p>Mark, put html here!</p>
</body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/", handler)
		err := http.ListenAndServe(":6123", nil)

		if err != nil {
			log.Println(err)
				os.Exit(1)
		}
}
