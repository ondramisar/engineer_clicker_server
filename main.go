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
	"time"
)

var userLink = "User"
var userMachineLink = "UsersMachines"
var userWorkerLink = "UserWorker"
var userMaterialLink = "UserMaterial"
var defaultMachineLink = "DefaultMachines"
var defaultWorkerLink = "DefaultWorker"
var defaultMaterialsLink = "DefaultMaterial"

func GetDefaultMachinesEndpoint(w http.ResponseWriter, req *http.Request) {
	var machines []DefaultMachine

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection(defaultMachineLink).Documents(ctx)
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
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection(userMachineLink).Where("idUser", "==", params["id"]).Documents(ctx)
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
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	dsnap, err := client.Collection(userMachineLink).Doc(person.ID).Set(ctx, person)
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
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc, err := client.Collection(userMachineLink).Doc(params["id"]).Set(ctx, map[string]interface{}{
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
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc, err := client.Collection(userMachineLink).Doc(params["id"]).Set(ctx, map[string]interface{}{
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
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection(defaultWorkerLink).Documents(ctx)
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
	var userWorkers []UserWorker
	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection(userWorkerLink).Where("idUser", "==", params["id"]).Documents(ctx)
	for {
		doc, err := doc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to: %v", err)
		}
		var c UserWorker
		doc.DataTo(&c)
		c.ID = doc.Ref.ID
		fmt.Printf("Document data: %#v\n", c)
		userWorkers = append(userWorkers, c)

	}
	json.NewEncoder(w).Encode(userWorkers)
}

func CreateWorker(w http.ResponseWriter, req *http.Request) {
	var person UserWorker
	_ = json.NewDecoder(req.Body).Decode(&person)

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	dsnap, err := client.Collection(userWorkerLink).Doc(person.ID).Set(ctx, person)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	_ = dsnap
	json.NewEncoder(w).Encode("succ")
}

func GetDefaultMaterialsEndpoint(w http.ResponseWriter, req *http.Request) {
	var material []DefaultMaterial

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection(defaultMaterialsLink).Documents(ctx)
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
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc := client.Collection(userMaterialLink).Where("idUser", "==", params["id"]).Documents(ctx)
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

func CreateMaterial(w http.ResponseWriter, req *http.Request) {
	var userMaterial UserMaterial
	_ = json.NewDecoder(req.Body).Decode(&userMaterial)

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	dsnap, err := client.Collection(userMaterialLink).Doc(userMaterial.ID).Set(ctx, userMaterial)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	_ = dsnap
	json.NewEncoder(w).Encode("succ")
}

func UpdateUserMaterialNumberOf(w http.ResponseWriter, req *http.Request) {
	var material UserMaterial
	_ = json.NewDecoder(req.Body).Decode(&material)

	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	materialUpdate, err := client.Collection(userMaterialLink).Doc(params["id"]).Set(ctx, map[string]interface{}{
		"numberOf": material.NumberOf,
	}, firestore.MergeAll)

	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	_ = materialUpdate

	json.NewEncoder(w).Encode("succ ")
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc, err := client.Collection(userLink).Doc(person.IdUser).Set(ctx, person)
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
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	doc, err := client.Collection(userLink).Doc(person.IdUser).Set(ctx, person)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	_ = doc
	json.NewEncoder(w).Encode("succ ")
}

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	dsnap, err := client.Collection(userLink).Doc(params["id"]).Get(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	var c User
	dsnap.DataTo(&c)
	c.IdUser = dsnap.Ref.ID
	fmt.Printf("Document data: %#v\n", c)

	json.NewEncoder(w).Encode(c)
}

func UpdateLastOutOfApp(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	userUpdate, err := client.Collection(userLink).Doc(params["id"]).Set(ctx, map[string]interface{}{
		"lastTimeOutOfApp": time.Now().UnixNano() / int64(time.Millisecond),
	}, firestore.MergeAll)

	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	_ = userUpdate

	json.NewEncoder(w).Encode("succ ")
}

func UpdateBackgroundUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	ctx := appengine.NewContext(req)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("fluffy-fox-project-firebase-adminsdk-wkhyq-b6739bc93c.json"))
	if err != nil {
		log.Fatalf("Failed to client: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}

	userSnap, err := client.Collection(userLink).Doc(params["id"]).Get(ctx)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	var user User
	userSnap.DataTo(&user)
	user.IdUser = userSnap.Ref.ID

	var userMachine []UserMachine
	userMachinesDoc := client.Collection(userMachineLink).Where("idUser", "==", params["id"]).Documents(ctx)
	for {
		doc, err := userMachinesDoc.Next()
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
		userMachine = append(userMachine, c)
	}

	timeNow := time.Now().UnixNano() / int64(time.Millisecond)
	timeUser := user.LastTimeOutOfApp
	difTime := (timeNow - timeUser) / 1000

	for _, mach := range userMachine {
		if mach.WorkerId != "" {
			machTime := difTime / mach.TimeToReach

			userMaterialDoc, err := client.Collection(userMaterialLink).Doc(mach.IdMaterialToGive).Get(ctx)
			if err != nil {
				log.Fatalf("Failed to: %v", err)
			}
			var userMaterial UserMaterial
			userMaterialDoc.DataTo(&userMaterial)
			userMaterial.IdUser = userMaterialDoc.Ref.ID

			userWorkerSnap, err := client.Collection(userWorkerLink).Doc(mach.WorkerId).Get(ctx)
			if err != nil {
				log.Fatalf("Failed to: %v", err)
			}
			var worker UserWorker
			userWorkerSnap.DataTo(&worker)
			worker.IdUser = userWorkerSnap.Ref.ID


			var numberToGive = userMaterial.NumberOf + (machTime * int64(mach.NumberOfMaterialsToGive*worker.MaterialMultiplayer))
			doc, err := client.Collection(userMaterialLink).Doc(mach.IdMaterialToGive).Set(ctx, map[string]interface{}{
				"numberOf": numberToGive,
			}, firestore.MergeAll)

			if err != nil {
				log.Fatalf("Failed to: %v", err)
			}
			_ = doc
		}
	}

	lastPayment := user.LastPayment

	dayToday := time.Now().Day()
	dayOfLastPayment := time.Unix(0, lastPayment*int64(time.Millisecond)).Day()
	difference := dayToday - dayOfLastPayment

	if difference >= 1 {
		var userWorkers []UserWorker
		userWorkersDoc := client.Collection(userWorkerLink).Where("idUser", "==", params["id"]).Documents(ctx)
		for {
			doc, err := userWorkersDoc.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to: %v", err)
			}
			var userWorker UserWorker
			doc.DataTo(&userWorker)
			userWorker.ID = doc.Ref.ID
			fmt.Printf("Document data: %#v\n", userWorker)
			userWorkers = append(userWorkers, userWorker)

		}

		for _, worker := range userWorkers {
			userUpdate, err := client.Collection(userLink).Doc(user.IdUser).Set(ctx, map[string]interface{}{
				"coins": user.Coins - (worker.Payment * difference),
			}, firestore.MergeAll)

			if err != nil {
				log.Fatalf("Failed to: %v", err)
			}
			_ = userUpdate
		}

		userUpdate, err := client.Collection(userLink).Doc(user.IdUser).Set(ctx, map[string]interface{}{
			"lastPayment": time.Now().UnixNano() / int64(time.Millisecond),
		}, firestore.MergeAll)

		if err != nil {
			log.Fatalf("Failed to: %v", err)
		}
		_ = userUpdate
	}

	json.NewEncoder(w).Encode("END")
}

func init() {
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/defaultMachines", http.StatusFound))
	router.HandleFunc("/updateBackgroundUser/{id}", UpdateBackgroundUser).Methods("POST")

	router.HandleFunc("/defaultMachines", GetDefaultMachinesEndpoint).Methods("GET")
	router.HandleFunc("/userMachines/{id}", GetUserMachinesEndpoint).Methods("GET")
	router.HandleFunc("/createMachine", CreateMachine).Methods("POST")
	router.HandleFunc("/addWorker/{id}/{id2}", AddWorkerToMachine).Methods("POST")
	router.HandleFunc("/removeWorker/{id}", RemoveWorkerToMachine).Methods("POST")

	router.HandleFunc("/defaultWorkers", GetDefaultWorkersEndpoint).Methods("GET")
	router.HandleFunc("/userWorkers/{id}", GetUserWorkersEndpoint).Methods("GET")
	router.HandleFunc("/createWorker", CreateWorker).Methods("POST")

	router.HandleFunc("/defaultMaterials", GetDefaultMaterialsEndpoint).Methods("GET")
	router.HandleFunc("/userMaterials/{id}", GetUserMaterialsEndpoint).Methods("GET")
	router.HandleFunc("/createMaterial", CreateMaterial).Methods("POST")
	router.HandleFunc("/updateMaterialNumberOf/{id}", UpdateUserMaterialNumberOf).Methods("POST")

	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/updateUser/{id}", UpdateUser).Methods("POST")
	router.HandleFunc("/lastOutOfApp/{id}", UpdateLastOutOfApp).Methods("POST")

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
}
