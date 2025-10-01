package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/scailo/go-sdk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var serverURL string
var envUsername string
var envPassword string

func getServerURL() string {
	return serverURL
}

func getLoginCreds() (username string, password string) {
	return envUsername, envPassword
}

var dir string

func main() {

	serverURLPtr := flag.String("server_url", "127.0.0.1:21000", "Enter the server URL.")

	envFilePathPtr := flag.String("env-file", "./.env", "Enter the path to the .env file.")

	modePtr := flag.String("mode", "", "Enter the operation mode: download or upload. 'download' would download all the templates from the server to a 'templates' folder. 'upload' would upload all the templates from the 'uploads' folder to the server.")
	dirPtr := flag.String("dir", "templates", "Enter the folder where the CSV files will be downloaded in case mode is download, or the folder where the CSV files will be uploaded from in case mode is upload.")

	shouldVerifyPtr := flag.Bool("should-verify", false, "Enter true if the user would like to verify all the added records as well. Default is false. Applicable only when mode is 'upload'.")
	shouldApprovePtr := flag.Bool("should-approve", false, "Enter true if the user would like to approve all the added records as well. Default is false. Applicable only when mode is 'upload'.")

	flag.Parse()

	if *modePtr != "upload" && *modePtr != "download" {
		flag.Usage()
		// fmt.Println("Invalid mode: `", *modePtr, "`. Please enter `download` or `upload`. Execute the command with the flag: --mode: `upload` or --mode: `download`")
		os.Exit(1)
	}

	godotenv.Load(*envFilePathPtr)

	serverURL = os.Getenv("server_url")
	if serverURL == "" {
		serverURL = *serverURLPtr
	}

	envUsername = os.Getenv("username")
	envPassword = os.Getenv("password")

	if len(envUsername) == 0 {
		fmt.Println("Username is empty. Please enter the username in the .env file.")
		os.Exit(1)
	}
	if len(envPassword) == 0 {
		fmt.Println("Password is empty. Please enter the password in the .env file.")
		os.Exit(1)
	}

	fmt.Println("Server URL is: ", getServerURL())
	fmt.Println("Username is: ", envUsername)
	fmt.Println("Mode is: ", *modePtr)
	fmt.Println("Dir is: ", *dirPtr)

	conn, err := grpc.NewClient(getServerURL(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()

	loginClient := sdk.NewLoginServiceClient(conn)
	username, password := getLoginCreds()
	loginResp, err := loginClient.LoginAsEmployeePrimary(ctx, &sdk.UserLoginRequest{
		Username:          username,
		PlainTextPassword: password,
	})
	if err != nil {
		panic(err)
	}

	md := metadata.Pairs(
		"auth_token", loginResp.AuthToken,
	)

	// 4. Create a new context with the metadata attached.
	ctx = metadata.NewOutgoingContext(ctx, md)

	dir = *dirPtr
	if *modePtr == "download" {
		err := os.RemoveAll(dir)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}

		downloadPrimaryRecords(ctx, conn)
		downloadLineItems(ctx, conn)
	} else if *modePtr == "upload" {
		uploadPrimaryRecords(ctx, conn, *shouldVerifyPtr, *shouldApprovePtr)
		uploadLineItems(ctx, conn)
	}
}

func getCSVFileNameForPrimaryRecords(f *sdk.StandardFile) string {
	return strings.Split(f.Name, ".")[0] + ".csv"
}

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	return records, err
}

func savePrimaryRecordsToFile(f *sdk.StandardFile, err error) {
	if err != nil {
		panic(err)
	}

	fileName := getCSVFileNameForPrimaryRecords(f)
	file, err := os.Create(path.Join(dir, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(f.Content)
	if err != nil {
		panic(err)
	}
	err = file.Sync()
	if err != nil {
		panic(err)
	}
}

func getCSVFileNameForLineItems(f *sdk.StandardFile) string {
	return strings.Split(f.Name, ".")[0] + ".csv"
}

func saveLineItemsToFile(f *sdk.StandardFile, err error) {
	if err != nil {
		panic(err)
	}

	fileName := getCSVFileNameForLineItems(f)
	file, err := os.Create(path.Join(dir, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(f.Content)
	if err != nil {
		panic(err)
	}
	err = file.Sync()
	if err != nil {
		panic(err)
	}
}
