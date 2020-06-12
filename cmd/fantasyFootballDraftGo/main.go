package main

import (
	"fmt"
	"os"

	"github.com/fantasyFootballDraftGo/internal/db/migrate"
	"github.com/fantasyFootballDraftGo/internal/restapi"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createDraftBoard(leagueName string) string {
	draftBoard := fmt.Sprintf("../../pkg/draftBoards/%s.db", leagueName)
	if !fileExists(draftBoard) {
		fmt.Println("Creating draft board...")
		gdb, err := gorm.Open("sqlite3", draftBoard)
		if err != nil {
			fmt.Printf("Error opening DB: %s", err)
		}

		migrate := migrate.New(gdb)
		migrate.Migrate()
		gdb.Close()
		return draftBoard
	}
	fmt.Println("Draft board already created...")
	return draftBoard
}

func main() {
	fmt.Print("Enter League Name: ")
	var leagueName string
	fmt.Scanln(&leagueName)
	draftBoard := createDraftBoard(leagueName)

	db, err := gorm.Open("sqlite3", draftBoard)
	if err != nil {
		fmt.Printf("Error opening DB: %s", err)
	}

	api := restapi.New(db)
	fmt.Printf("Starting Draft for %s...\n", leagueName)
	api.HandleRequests()
}
