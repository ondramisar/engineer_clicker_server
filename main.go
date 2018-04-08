package main

import (
	"encoding/json"
	"net/http"
	"os"

	"gorilla/mux"
	"gorilla/handlers"
	"google.golang.org/appengine"
	"fmt"

	"log"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/api/iterator"
	"cloud.google.com/go/firestore"
)

func GetDefaultMachinesEndpoint(w http.ResponseWriter, req *http.Request) {
	var machines []DefaultMachine

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection("DefaultMachines").Documents(ctx)
	for {
		doc, err := doc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to: %v", err)
		}
		var c DefaultMachine
		doc.DataTo(&c)
		c.ID = doc.Ref.ID
		fmt.Printf("Document data: %#v\n", c)
		machines = append(machines, c)

	}
	json.NewEncoder(w).Encode(machines)
}

func GetUserMachinesEndpoint(w http.ResponseWriter, req *http.Request) {
	var userMachines []UserMachine
	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection("UsersMachines").Where("idUser", "==", params["id"]).Documents(ctx)
	for {
		doc, err := doc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to: %v", err)
		}
		var c UserMachine
		doc.DataTo(&c)
		c.ID = doc.Ref.ID
		fmt.Printf("Document data: %#v\n", c)
		userMachines = append(userMachines, c)

	}
	json.NewEncoder(w).Encode(userMachines)
}

func CreateMachine(w http.ResponseWriter, req *http.Request) {

	var person UserMachine
	_ = json.NewDecoder(req.Body).Decode(&person)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	dsnap, err := client.Collection("UsersMachines").Doc(person.ID).Set(ctx, person)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	_ = dsnap
	json.NewEncoder(w).Encode("succ")
}

func AddWorkerToMachine(w http.ResponseWriter, req *http.Request) {
	var id string
	_ = json.NewDecoder(req.Body).Decode(id)

	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc, err := client.Collection("UsersMachines").Doc(params["id"]).Set(ctx, map[string]interface{}{
		"workerId": params["id2"],
	}, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	_ = doc
	json.NewEncoder(w).Encode("succ")
}

func RemoveWorkerToMachine(w http.ResponseWriter, req *http.Request) {
	var id string
	_ = json.NewDecoder(req.Body).Decode(id)

	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc, err := client.Collection("UsersMachines").Doc(params["id"]).Set(ctx, map[string]interface{}{
		"workerId": "",
	}, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	_ = doc
	json.NewEncoder(w).Encode("succ")
}

func GetDefaultWorkersEndpoint(w http.ResponseWriter, req *http.Request) {

	var workers []DefaultWorker

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection("DefaultWorker").Documents(ctx)
	for {
		doc, err := doc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to: %v", err)
		}
		var c DefaultWorker
		doc.DataTo(&c)
		c.ID = doc.Ref.ID
		fmt.Printf("Document data: %#v\n", c)
		workers = append(workers, c)

	}
	json.NewEncoder(w).Encode(workers)
}

func GetUserWorkersEndpoint(w http.ResponseWriter, req *http.Request) {

	var userWorkers []UsertWorker
	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection("UserWorker").Where("idUser", "==", params["id"]).Documents(ctx)
	for {
		doc, err := doc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to: %v", err)
		}
		var c UsertWorker
		doc.DataTo(&c)
		c.ID = doc.Ref.ID
		fmt.Printf("Document data: %#v\n", c)
		userWorkers = append(userWorkers, c)

	}
	json.NewEncoder(w).Encode(userWorkers)
}

func CreateWorker(w http.ResponseWriter, req *http.Request) {

	var person UsertWorker
	_ = json.NewDecoder(req.Body).Decode(&person)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	dsnap, err := client.Collection("UserWorker").Doc(person.ID).Set(ctx, person)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	_ = dsnap
	json.NewEncoder(w).Encode("succ")
}

func GetDefaultMaterialsEndpoint(w http.ResponseWriter, req *http.Request) {

	var material []DefaultMaterial

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection("DefaultMaterial").Documents(ctx)
	for {
		doc, err := doc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to: %v", err)
		}
		var c DefaultMaterial
		doc.DataTo(&c)
		c.ID = doc.Ref.ID
		fmt.Printf("Document data: %#v\n", c)
		material = append(material, c)

	}
	json.NewEncoder(w).Encode(material)
}

func GetUserMaterialsEndpoint(w http.ResponseWriter, req *http.Request) {

	var material []UserMaterial
	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection("UserMaterial").Where("idUser", "==", params["id"]).Documents(ctx)
	for {
		doc, err := doc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to: %v", err)
		}
		var c UserMaterial
		doc.DataTo(&c)
		c.ID = doc.Ref.ID
		fmt.Printf("Document data: %#v\n", c)
		material = append(material, c)

	}
	json.NewEncoder(w).Encode(material)
}

func GetUsersEndpoint(w http.ResponseWriter, req *http.Request) {

	json.NewEncoder(w).Encode("succ")
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc, err := client.Collection("User").Doc(person.IdUser).Set(ctx, person)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	_ = doc
	json.NewEncoder(w).Encode("succ ")
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc, err := client.Collection("User").Doc(person.IdUser).Set(ctx, person)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	_ = doc
	json.NewEncoder(w).Encode("succ ")
}

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	//config := &firebase.Config{ProjectID: "fluffy-fox-project"}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	dsnap, err := client.Collection("User").Doc(params["id"]).Get(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	var c User
	dsnap.DataTo(&c)
	c.IdUser = dsnap.Ref.ID
	fmt.Printf("Document data: %#v\n", c)

	json.NewEncoder(w).Encode(c)
}

func init() {

	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/defaultMachines", http.StatusFound))

	router.HandleFunc("/defaultMachines", GetDefaultMachinesEndpoint).Methods("GET")
	router.HandleFunc("/userMachines/{id}", GetUserMachinesEndpoint).Methods("GET")
	router.HandleFunc("/createMachine", CreateMachine).Methods("POST")
	router.HandleFunc("/addWorker/{id}/{id2}", AddWorkerToMachine).Methods("POST")
	router.HandleFunc("/removeWorker/{id}", AddWorkerToMachine).Methods("POST")

	router.HandleFunc("/defaultWorkers", GetDefaultWorkersEndpoint).Methods("GET")
	router.HandleFunc("/userWorkers/{id}", GetUserWorkersEndpoint).Methods("GET")
	router.HandleFunc("/createWorker", CreateWorker).Methods("POST")

	router.HandleFunc("/defaultMaterials", GetDefaultMaterialsEndpoint).Methods("GET")
	router.HandleFunc("/userMaterials/{id}", GetUserMaterialsEndpoint).Methods("GET")

	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/updateUser/{id}", UpdateUser).Methods("POST")
	router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")
	//log.Fatal(http.ListenAndServe(":3000", router))
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
}
