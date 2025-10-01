package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/scailo/go-sdk"
	"google.golang.org/grpc"
)

func getCSVRowsFromPrimaryRecordsFile(f *sdk.StandardFile, err error) (*sdk.StandardFile, error) {
	if err != nil {
		panic(err)
	}
	fName := getCSVFileNameForPrimaryRecords(f)
	csvFilePath := path.Join(dir, fName)

	rows, err := readCsvFile(csvFilePath)
	if len(rows) > 1 {
		content, err := os.ReadFile(csvFilePath)
		if err != nil {
			panic(err)
		}
		return &sdk.StandardFile{
			Name:     fName,
			MimeType: "text/csv",
			Content:  content,
		}, err
	}
	return nil, err
}

func getCSVRowsFromLineItemsFile(f *sdk.StandardFile, err error) (map[string]*sdk.StandardFile, error) {
	if err != nil {
		panic(err)
	}
	templateName := getCSVFileNameForLineItems(f)
	templatePrefix := strings.ReplaceAll(templateName, ".csv", "")

	fmt.Println("Processing template: ", templateName)

	var mapOfFiles = make(map[string]*sdk.StandardFile)

	// Identify all the files that match the template format
	matches, _ := filepath.Glob(dir + string(os.PathSeparator) + templatePrefix + "*.csv")
	for _, matchedFile := range matches {
		fileNameWithoutDir := strings.ReplaceAll(matchedFile, dir+string(os.PathSeparator), "")
		// Identify the identifier (the part of the file that will be used to identify the parent record)
		fileIdentifier := strings.ReplaceAll(strings.ReplaceAll(fileNameWithoutDir, templatePrefix+".", ""), ".csv", "")
		fmt.Println("Identifier is: ", fileIdentifier, ", File is: ", fileNameWithoutDir)

		csvFilePath := path.Join(dir, fileNameWithoutDir)
		rows, err := readCsvFile(csvFilePath)
		if err != nil {
			panic(err)
		}
		if len(rows) > 1 {
			content, err := os.ReadFile(csvFilePath)
			if err != nil {
				panic(err)
			}
			mapOfFiles[fileIdentifier] = &sdk.StandardFile{
				Name:     fileNameWithoutDir,
				MimeType: "text/csv",
				Content:  content,
			}
		}
	}
	return mapOfFiles, err
}

func initAddedList() *sdk.IdentifierUUIDsList {
	var addedList = new(sdk.IdentifierUUIDsList)
	addedList.List = make([]*sdk.IdentifierUUID, 0)
	return addedList
}

func uploadActionCodesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadActionCodesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewActionsCodesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading action codes: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded action codes from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending action code for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying action code is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving action code is: %v", err))
				}
			}
		}
	}
}

func uploadActivitiesGroupsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadActivitiesGroupsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewActivitiesGroupsServiceClient(conn)

	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading activity groups: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded activity groups from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending activity group for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying activity group is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving activity group is: %v", err))
				}
			}
		}
	}
}

func uploadActivitiesStatusesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadActivitiesStatusesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewActivitiesStatusesServiceClient(conn)

	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading activity statuses: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded activity statuses from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending activity status for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying activity status is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving activity status is: %v", err))
				}
			}
		}
	}
}

func uploadActivitiesTagsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadActivitiesTagsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewActivitiesTagsServiceClient(conn)

	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading activity tags: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded activity tags from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending activity tag for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying activity tag is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving activity tag is: %v", err))
				}
			}
		}
	}
}

func uploadActivitiesTemplate(ctx context.Context, conn *grpc.ClientConn) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadActivitiesTemplate(ctx, conn))
	if toUploadFile != nil {
		sdk.NewActivitiesServiceClient(conn).ImportFromCSV(ctx, toUploadFile)
	}
}

func uploadAnnouncementsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadAnnouncementsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewAnnouncementsServiceClient(conn)

	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading announcements: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded announcements from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending announcement for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying announcement is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving announcement is: %v", err))
				}
			}
		}
	}
}

func uploadAssociatesTemplate(ctx context.Context, conn *grpc.ClientConn) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadAssociatesTemplate(ctx, conn))
	if toUploadFile != nil {
		sdk.NewAssociatesServiceClient(conn).ImportFromCSV(ctx, toUploadFile)
	}
}

func uploadBankAccountsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadBankAccountsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewBankAccountsServiceClient(conn)

	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading bank accounts: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded bank accounts from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending bank account for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying bank account is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving bank account is: %v", err))
				}
			}
		}
	}
}

func uploadClientsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadClientsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewClientsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading clients: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded clients from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending client for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying client is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving client is: %v", err))
				}
			}
		}
	}
}

func uploadComponentsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldSendToStore bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadComponentsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewComponentsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading components: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded components from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldSendToStore {
				record, err := c.ViewByUUID(ctx, &sdk.IdentifierUUID{
					Uuid: identifier.Uuid,
				})
				if err != nil {
					panic(fmt.Errorf("error while retrieving the uploaded component is: %v", err))
				}

				_, err = c.SendToStore(ctx, &sdk.ComponentsServiceSendToStoreRequest{
					Id:          record.Metadata.Id,
					UserComment: "Sending to store from uploader program",
					StoreId:     1,
				})
				if err != nil {
					panic(fmt.Errorf("error while sending component to store is: %v", err))
				}
			}
		}
	}
}

func uploadCurrenciesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadCurrenciesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewCurrenciesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading currencies: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded currencies from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending currency for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying currency is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving currency is: %v", err))
				}
			}
		}
	}
}

func uploadDepartmentsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadDepartmentsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewDepartmentsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading departments: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded departments from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending department for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying department is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving department is: %v", err))
				}
			}
		}
	}
}

func uploadEquationsFamiliesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadEquationsFamiliesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewEquationsFamiliesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading equations families: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded equations families from: %s\n", toUploadFile.Name)
	}

	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Equations Families Line Items template")

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadEquationsFamiliesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByName(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadEquationFamilyItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending equation family for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying equation family is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving equation family is: %v", err))
				}
			}
		}
	}
}

func uploadEquationsReplaceablesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadEquationsReplaceablesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewEquationsReplaceablesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading equations replaceables: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded equations replaceables from: %s\n", toUploadFile.Name)
	}

	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Equations Replaceables Line Items template")

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadEquationsReplaceablesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByName(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadEquationReplaceableItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending equation replaceable for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying equation replaceable is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving equation replaceable is: %v", err))
				}
			}
		}
	}
}

func uploadEquationsSalesBundlesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadEquationsSalesBundlesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewEquationsSalesBundlesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading equations sales bundles: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded equations sales bundles from: %s\n", toUploadFile.Name)
	}

	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Equations Sales Bundles Line Items template")

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadEquationsSalesBundlesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByName(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadEquationSalesBundleItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending equation sales bundle for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying equation sales bundle is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving equation sales bundle is: %v", err))
				}
			}
		}
	}
}

func uploadEquationsWorkOrdersTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadEquationsWorkOrdersTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewEquationsWorkOrdersServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading equations work orders: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded equations work orders from: %s\n", toUploadFile.Name)
	}

	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Equations Work Orders Line Items template")

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadEquationsWorkOrdersLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByName(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadEquationWorkOrderItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending equation work order for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying equation work order is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving equation work order is: %v", err))
				}
			}
		}
	}
}

func uploadEquipmentsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldSendToStore bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadEquipmentsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewEquipmentsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading equipments: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded equipments from: %s\n", toUploadFile.Name)
	}

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldSendToStore {
				record, err := c.ViewByUUID(ctx, &sdk.IdentifierUUID{
					Uuid: identifier.Uuid,
				})
				if err != nil {
					panic(fmt.Errorf("error while retrieving the uploaded equipment is: %v", err))
				}

				_, err = c.SendToStore(ctx, &sdk.EquipmentsServiceSendToStoreRequest{
					Id:          record.Metadata.Id,
					UserComment: "Sending to store from uploader program",
					StoreId:     1,
				})
				if err != nil {
					panic(fmt.Errorf("error while sending equipment to store is: %v", err))
				}
			}
		}
	}
}

func uploadFamiliesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadFamiliesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewFamiliesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading families: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded families from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending family for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying family is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving family is: %v", err))
				}
			}
		}
	}
}

func uploadFeedstocksTemplate(ctx context.Context, conn *grpc.ClientConn, shouldSendToStore bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadFeedstocksTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewFeedstocksServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading feedstocks: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded feedstocks from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldSendToStore {
				record, err := c.ViewByUUID(ctx, &sdk.IdentifierUUID{
					Uuid: identifier.Uuid,
				})
				if err != nil {
					panic(fmt.Errorf("error while retrieving the uploaded feedstock is: %v", err))
				}

				_, err = c.SendToStore(ctx, &sdk.FeedstocksServiceSendToStoreRequest{
					Id:          record.Metadata.Id,
					UserComment: "Sending to store from uploader program",
					StoreId:     1,
				})
				if err != nil {
					panic(fmt.Errorf("error while sending feedstock to store is: %v", err))
				}
			}
		}
	}
}

func uploadFormsFieldsTemplate(ctx context.Context, conn *grpc.ClientConn) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadFormsFieldsTemplate(ctx, conn))
	if toUploadFile != nil {
		sdk.NewFormsFieldsServiceClient(conn).ImportFromCSV(ctx, toUploadFile)
	}
}

func uploadFormsSectionsTemplate(ctx context.Context, conn *grpc.ClientConn) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadFormsSectionsTemplate(ctx, conn))
	if toUploadFile != nil {
		sdk.NewFormsSectionsServiceClient(conn).ImportFromCSV(ctx, toUploadFile)
	}
}

func uploadHolidaysTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadHolidaysTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewHolidaysServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading holidays: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded holidays from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending holiday for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying holiday is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving holiday is: %v", err))
				}
			}
		}
	}
}

func uploadInfrastructuresTemplate(ctx context.Context, conn *grpc.ClientConn, shouldSendToStore bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadInfrastructuresTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewInfrastructuresServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading infrastructures: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded infrastructures from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldSendToStore {
				record, err := c.ViewByUUID(ctx, &sdk.IdentifierUUID{
					Uuid: identifier.Uuid,
				})
				if err != nil {
					panic(fmt.Errorf("error while retrieving the uploaded infrastructure is: %v", err))
				}

				_, err = c.SendToStore(ctx, &sdk.InfrastructuresServiceSendToStoreRequest{
					Id:          record.Metadata.Id,
					UserComment: "Sending to store from uploader program",
					StoreId:     1,
				})
				if err != nil {
					panic(fmt.Errorf("error while sending infrastructure to store is: %v", err))
				}
			}
		}
	}
}

func uploadLabelsTemplate(ctx context.Context, conn *grpc.ClientConn) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadLabelsTemplate(ctx, conn))
	if toUploadFile != nil {
		sdk.NewLabelsServiceClient(conn).ImportFromCSV(ctx, toUploadFile)
	}
}

func uploadLeavesTypesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadLeavesTypesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewLeavesTypesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading leave types: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded leave types from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending leave type for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying leave type is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving leave type is: %v", err))
				}
			}
		}
	}
}

func uploadLedgersTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadLedgersTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewLedgersServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading ledgers: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded ledgers from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending ledger for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying ledger is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving ledger is: %v", err))
				}
			}
		}
	}
}

func uploadLocationsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadLocationsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewLocationsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading locations: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded locations from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending location for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying location is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving location is: %v", err))
				}
			}
		}
	}
}

func uploadMerchandisesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldSendToStore bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadMerchandisesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewMerchandisesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading merchandises: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded merchandises from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldSendToStore {
				record, err := c.ViewByUUID(ctx, &sdk.IdentifierUUID{
					Uuid: identifier.Uuid,
				})
				if err != nil {
					panic(fmt.Errorf("error while retrieving the uploaded merchandise is: %v", err))
				}

				_, err = c.SendToStore(ctx, &sdk.MerchandisesServiceSendToStoreRequest{
					Id:          record.Metadata.Id,
					UserComment: "Sending to store from uploader program",
					StoreId:     1,
				})
				if err != nil {
					panic(fmt.Errorf("error while sending merchandise to store is: %v", err))
				}
			}
		}
	}
}

func uploadPayrollGroupsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadPayrollGroupsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewPayrollGroupsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading payroll groups: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded payroll groups from: %s\n", toUploadFile.Name)
	}

	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Payroll Groups Line Items template")

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadPayrollGroupsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByCode(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadPayrollGroupItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending payroll group for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying payroll group is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving payroll group is: %v", err))
				}
			}
		}
	}
}

func uploadPayrollParamsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadPayrollParamsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewPayrollParamsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading payroll params: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded payroll params from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending payroll param for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying payroll param is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving payroll param is: %v", err))
				}
			}
		}
	}
}

func uploadProductsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldSendToStore bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadProductsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewProductsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading products: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded products from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldSendToStore {
				record, err := c.ViewByUUID(ctx, &sdk.IdentifierUUID{
					Uuid: identifier.Uuid,
				})
				if err != nil {
					panic(fmt.Errorf("error while retrieving the uploaded product is: %v", err))
				}

				_, err = c.SendToStore(ctx, &sdk.ProductsServiceSendToStoreRequest{
					Id:          record.Metadata.Id,
					UserComment: "Sending to store from uploader program",
					StoreId:     1,
				})
				if err != nil {
					panic(fmt.Errorf("error while sending product to store is: %v", err))
				}
			}
		}
	}
}

func uploadQCGroupsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadQCGroupsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewQCGroupsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading qc groups: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded qc groups from: %s\n", toUploadFile.Name)
	}

	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload QC Groups Line Items template")

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadQCGroupsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByCode(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadQCGroupItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending qc group for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying qc group is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving qc group is: %v", err))
				}
			}
		}
	}
}

func uploadQCParamsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadQCParamsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewQCParamsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading qc params: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded qc params from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending qc param for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying qc param is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving qc param is: %v", err))
				}
			}
		}
	}
}

func uploadRolesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadRolesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewRolesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading roles: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded roles from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending role for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying role is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving role is: %v", err))
				}
			}
		}
	}
}

func uploadShiftsGroupsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadShiftsGroupsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewShiftsGroupsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading shift groups: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded shift groups from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending shift group for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying shift group is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving shift group is: %v", err))
				}
			}
		}
	}
}

func uploadShiftsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadShiftsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewShiftsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading shifts: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded shifts from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending shift for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying shift is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving shift is: %v", err))
				}
			}
		}
	}
}

func uploadSkillsGroupsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadSkillsGroupsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewSkillsGroupsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading skill groups: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded skill groups from: %s\n", toUploadFile.Name)
	}

	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Skill Groups Line Items template")

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadSkillGroupsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByCode(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadSkillGroupItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending skill group for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying skill group is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving skill group is: %v", err))
				}
			}
		}
	}
}

func uploadSkillsParamsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadSkillsParamsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewSkillsParamsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading skill params: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded skill params from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending skill param for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying skill param is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving skill param is: %v", err))
				}
			}
		}
	}
}

func uploadStoragesTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadStoragesTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewStoragesServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading storages: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded storages from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending storage for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying storage is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving storage is: %v", err))
				}
			}
		}
	}
}

func uploadStoresTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadStoresTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewStoresServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading stores: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded stores from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending store for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying store is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving store is: %v", err))
				}
			}
		}
	}
}

func uploadTaxGroupsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadTaxGroupsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewTaxGroupsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading tax groups: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded tax groups from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending tax group for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying tax group is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving tax group is: %v", err))
				}
			}
		}
	}
}

func uploadTaxParamsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadTaxParamsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewTaxParamsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading tax params: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded tax params from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending tax param for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying tax param is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving tax param is: %v", err))
				}
			}
		}
	}
}

func uploadTeamsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadTeamsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewTeamsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading teams: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded teams from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending team for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying team is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving team is: %v", err))
				}
			}
		}
	}
}

func uploadUnitsOfMaterialsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadUnitsOfMaterialsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewUnitsOfMaterialsServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading units of materials: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded units of materials from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending units of material for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying units of material is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving units of material is: %v", err))
				}
			}
		}
	}
}

func uploadUsersTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadUsersTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewUsersServiceClient(conn)
	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading users: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded users from: %s\n", toUploadFile.Name)
	}
	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending user for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying user is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving user is: %v", err))
				}
			}
		}
	}
}

func uploadVendorsTemplate(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	toUploadFile, _ := getCSVRowsFromPrimaryRecordsFile(downloadVendorsTemplate(ctx, conn))
	var addedList = initAddedList()
	var err error
	c := sdk.NewVendorsServiceClient(conn)

	if toUploadFile != nil {
		addedList, err = c.ImportFromCSV(ctx, toUploadFile)
		if err != nil {
			panic(fmt.Errorf("error while uploading vendors: %v from file: %s", err, toUploadFile.Name))
		}
		fmt.Printf("successfully uploaded vendors from: %s\n", toUploadFile.Name)
	}

	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Vendors Line Items template")

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadVendorsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByCode(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadVendorItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------

	if len(addedList.List) > 0 {
		for _, identifier := range addedList.List {
			if shouldVerify {
				_, err = c.SendForVerification(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Sending for verification from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while sending vendor for verification is: %v", err))
				}
				_, err = c.Verify(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Verifying from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while verifying vendor is: %v", err))
				}
			}
			if shouldVerify && shouldApprove {
				_, err = c.Approve(ctx, &sdk.IdentifierUUIDWithUserComment{
					Uuid:        identifier.Uuid,
					UserComment: "Approving from uploader program",
				})
				if err != nil {
					panic(fmt.Errorf("error while approving vendor is: %v", err))
				}
			}
		}
	}
}

func uploadPrimaryRecords(ctx context.Context, conn *grpc.ClientConn, shouldVerify bool, shouldApprove bool) {
	uploadFormsSectionsTemplate(ctx, conn)
	uploadFormsFieldsTemplate(ctx, conn)

	uploadActionCodesTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadActivitiesGroupsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadActivitiesStatusesTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadActivitiesTagsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadActivitiesTemplate(ctx, conn)

	uploadAnnouncementsTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadLocationsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadBankAccountsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadCurrenciesTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadClientsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadUnitsOfMaterialsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadLedgersTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadLabelsTemplate(ctx, conn)

	uploadRolesTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadUsersTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadTeamsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadDepartmentsTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadStoresTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadStoragesTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadTaxParamsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadTaxGroupsTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadQCParamsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadQCGroupsTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadFamiliesTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadVendorsTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadComponentsTemplate(ctx, conn, shouldApprove)
	uploadInfrastructuresTemplate(ctx, conn, shouldApprove)
	uploadMerchandisesTemplate(ctx, conn, shouldApprove)
	uploadEquipmentsTemplate(ctx, conn, shouldApprove)
	uploadFeedstocksTemplate(ctx, conn, shouldApprove)
	uploadProductsTemplate(ctx, conn, shouldApprove)

	uploadShiftsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadShiftsGroupsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadHolidaysTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadPayrollParamsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadPayrollGroupsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadLeavesTypesTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadSkillsParamsTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadSkillsGroupsTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadEquationsFamiliesTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadEquationsReplaceablesTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadEquationsSalesBundlesTemplate(ctx, conn, shouldVerify, shouldApprove)
	uploadEquationsWorkOrdersTemplate(ctx, conn, shouldVerify, shouldApprove)

	uploadAssociatesTemplate(ctx, conn)

	fmt.Println("Uploaded all relevant files")
}

func uploadAssetIndentsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Asset Indents Line Items template")
	c := sdk.NewAssetIndentsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadAssetIndentsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadAssetIndentItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadCreditNotesLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Credit Notes Line Items template")
	c := sdk.NewCreditNotesServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadCreditNotesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadCreditNoteItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadDebitNotesLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Debit Notes Line Items template")
	c := sdk.NewDebitNotesServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadDebitNotesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadDebitNoteItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadExpensesLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Expenses Line Items template")
	c := sdk.NewExpensesServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadExpensesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadExpenseItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadGoalsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Goals Line Items template")
	c := sdk.NewGoalsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadGoalsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadGoalItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadGoodsReceiptsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Goods Receipts Line Items template")
	c := sdk.NewGoodsReceiptsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadGoodsReceiptsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadGoodsReceiptItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadProductionIndentsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Production Indents Line Items template")
	c := sdk.NewProductionIndentsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadProductionIndentsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadProductionIndentItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadProductionPlansLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Production Plans Line Items template")
	c := sdk.NewProductionPlansServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadProductionPlansLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadProductionPlanItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadProformaInvoicesLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Proforma Invoices Line Items template")
	c := sdk.NewProformaInvoicesServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadProformaInvoicesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadProformaInvoiceItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadPurchaseEnquiriesLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Purchase Enquiries Line Items template")
	c := sdk.NewPurchasesEnquiriesServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadPurchaseEnquiriesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadPurchaseEnquiryItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadPurchaseIndentsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Purchase Indents Line Items template")
	c := sdk.NewPurchasesIndentsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadPurchaseIndentsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadPurchaseIndentItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadPurchaseOrdersLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Purchase Orders Line Items template")
	c := sdk.NewPurchasesOrdersServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadPurchaseOrdersLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadPurchaseOrderItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadQuotationRequestsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Quotation Requests Line Items template")
	c := sdk.NewQuotationsRequestsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadQuotationRequestsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadQuotationRequestItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadReplaceableIndentsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Replaceable Indents Line Items template")
	c := sdk.NewReplaceableIndentsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadReplaceableIndentsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadReplaceableIndentItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadSalesEnquiriesLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Sales Enquiries Line Items template")
	c := sdk.NewSalesEnquiriesServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadSalesEnquiriesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadSalesEnquiryItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadSalesInvoicesLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Sales Invoices Line Items template")
	c := sdk.NewSalesInvoicesServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadSalesInvoicesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadSalesInvoiceItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadSalesOrdersLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Sales Orders Line Items template")
	c := sdk.NewSalesOrdersServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadSalesOrdersLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadSalesOrderItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadSalesQuotationsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Sales Quotations Line Items template")
	c := sdk.NewSalesQuotationsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadSalesQuotationsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadSalesQuotationItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadStockAuditsLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Stock Audits Line Items template")
	c := sdk.NewStockAuditsServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadStockAuditsLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadStockAuditItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadSupplyOffersLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Supply Offers Line Items template")
	c := sdk.NewSupplyOffersServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadSupplyOffersLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadSupplyOfferItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadVendorInvoicesLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Vendor Invoices Line Items template")
	c := sdk.NewVendorInvoicesServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadVendorInvoicesLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadVendorInvoiceItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadWorkOrdersLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// ---------------------------------------------------------------------------------------------------------
	fmt.Println("About to upload Work Orders Line Items template")
	c := sdk.NewWorkOrdersServiceClient(conn)

	// Download the latest template and identify the file name
	toUploadFiles, _ := getCSVRowsFromLineItemsFile(downloadWorkOrdersLineItemsTemplate(ctx, conn))

	for code, f := range toUploadFiles {
		parentRecord, err := c.ViewByReferenceID(ctx, &sdk.SimpleSearchReq{
			SearchKey: code,
		})
		if err != nil {
			panic(err)
		}

		_, err = c.UploadWorkOrderItems(ctx, &sdk.IdentifierUUIDWithFile{
			Uuid:        parentRecord.Metadata.Uuid,
			FileContent: f.Content,
		})
		if err != nil {
			panic(err)
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func uploadLineItems(ctx context.Context, conn *grpc.ClientConn) {

	// Asset Indents
	uploadAssetIndentsLineItems(ctx, conn)
	// Credit Notes
	uploadCreditNotesLineItems(ctx, conn)
	// Debit Notes
	uploadDebitNotesLineItems(ctx, conn)

	// Expenses
	uploadExpensesLineItems(ctx, conn)
	// Goals
	uploadGoalsLineItems(ctx, conn)
	// Goods Receipts
	uploadGoodsReceiptsLineItems(ctx, conn)

	// Production Indents
	uploadProductionIndentsLineItems(ctx, conn)
	// Production Plans
	uploadProductionPlansLineItems(ctx, conn)
	// Proforma Invoices
	uploadProformaInvoicesLineItems(ctx, conn)
	// Purchase Enquiries
	uploadPurchaseEnquiriesLineItems(ctx, conn)
	// Purchase Indents
	uploadPurchaseIndentsLineItems(ctx, conn)
	// Purchase Orders
	uploadPurchaseOrdersLineItems(ctx, conn)

	// Quotation Requests
	uploadQuotationRequestsLineItems(ctx, conn)
	// Replaceable Indents
	uploadReplaceableIndentsLineItems(ctx, conn)
	// Sales Enquiries
	uploadSalesEnquiriesLineItems(ctx, conn)
	// Sales Invoices
	uploadSalesInvoicesLineItems(ctx, conn)
	// Sales Orders
	uploadSalesOrdersLineItems(ctx, conn)
	// Sales Quotations
	uploadSalesQuotationsLineItems(ctx, conn)

	// Stock Audits
	uploadStockAuditsLineItems(ctx, conn)
	// Supply Offers
	uploadSupplyOffersLineItems(ctx, conn)
	// Vendor Invoices
	uploadVendorInvoicesLineItems(ctx, conn)
	// Work Orders
	uploadWorkOrdersLineItems(ctx, conn)
}
