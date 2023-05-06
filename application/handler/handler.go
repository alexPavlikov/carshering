package handler

import "net/http"

func Handler() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../assets/"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/subscribe", subscribeHandler)

	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/reg", regHandler)

	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/product", productHandler)
	http.HandleFunc("/card", cardHandler)

	http.HandleFunc("/exit", exitHandler)
}
