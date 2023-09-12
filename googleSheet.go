package TelegramInviteRecorder

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewSheet(filepath, sheetName string) string {
	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentialsFile(filepath))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	sheet := sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: sheetName,
		},
	}

	spreadsheet, err := client.Spreadsheets.Create(&sheet).Do()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return spreadsheet.SpreadsheetId
}

func AddEditor(filepath, sheetID, gmail string) {

	ctx := context.Background()
	newService, err := drive.NewService(ctx, option.WithCredentialsFile(filepath))
	if err != nil {
		fmt.Println(err)
	}
	p := drive.Permission{
		AllowFileDiscovery:         false,
		Deleted:                    false,
		DisplayName:                "",
		Domain:                     "",
		EmailAddress:               gmail,
		ExpirationTime:             "",
		Id:                         "",
		Kind:                       "",
		PendingOwner:               false,
		PermissionDetails:          nil,
		PhotoLink:                  "",
		Role:                       "writer",
		TeamDrivePermissionDetails: nil,
		Type:                       "user",
		View:                       "",
		ServerResponse:             googleapi.ServerResponse{},
		ForceSendFields:            nil,
		NullFields:                 nil,
	}
	_, err = drive.NewPermissionsService(newService).Create(sheetID, &p).Do()
	if err != nil {
		fmt.Println(err)
	}
}
