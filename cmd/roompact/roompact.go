package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	fmt.Println("Hello, Roompact!")

	type userData struct {
		Email            string `json:"user_email"`
		First            string `json:"user_first"`
		Last             string `json:"user_last"`
		Phone            string `json:"user_phone"`
		Gender           string `json:"user_gender"`
		StudentID        string `json:"user_student_id"`
		Address          string `json:"user_address"`
		StartDate        string `json:"user_start_date"`
		EndDate          string `json:"user_end_date"`
		WelcomeEmailDate string `json:"user_welcome_email_date"`
		BuildingName     string `json:"building_name"`
		FloorName        string `json:"floor_name"`
		RoomName         string `json:"room_name"`
	}

	type request struct {
		Key      string     `json:"key"`
		Version  string     `json:"version"`
		UserData []userData `json:"user_data"`
	}

	gofakeit.Seed(0)

	req := request{
		Key:     "Elxob3Oiof1enbVNOxELygMMNafYI24nnCMyFAR4PTjc4f0TMFElWq7JBmHkpMHM",
		Version: "2.0",
	}
	var users []userData
	for i := 0; i < 10; i++ {
		p := gofakeit.Person()
		users = append(users, userData{
			Email:            gofakeit.Email(),
			First:            p.FirstName,
			Last:             p.LastName,
			Phone:            gofakeit.Phone(),
			Gender:           p.Gender[0:1],
			StudentID:        fmt.Sprintf("%d", gofakeit.Number(1111222200, 2111222200)),
			Address:          fmt.Sprintf("%v", gofakeit.Address()), // TODO: improve formatting
			StartDate:        gofakeit.Date().Format("2006-01-02"),
			EndDate:          gofakeit.Date().Format("2006-01-02"), // TODO: ensure after start date
			WelcomeEmailDate: gofakeit.Date().Format("2006-01-02"), // TODO: ensure before start date
			BuildingName:     gofakeit.PetName(),
			FloorName:        fmt.Sprintf("%d", gofakeit.Number(1, 10)), // TODO: format as 1st, 2nd, etc.
			RoomName:         strings.Title(gofakeit.HipsterWord()),
		})
	}
	req.UserData = users
	body, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, body, "", "\t")
	fmt.Println(string(out.String()))
	fmt.Println()

	resp, err := http.Post("https://roompact.com/api/users", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Indent(&out, respBody, "", "\t")
	fmt.Println(string(out.String()))
}
