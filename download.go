package main

import (
	"context"
	"fmt"

	"github.com/scailo/go-sdk"
	"google.golang.org/grpc"
)

func downloadActionCodesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewActionsCodesServiceClient(conn).DownloadAsCSV(ctx, &sdk.ActionsCodesServiceFilterReq{})
}

func downloadActivitiesGroupsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewActivitiesGroupsServiceClient(conn).DownloadAsCSV(ctx, &sdk.ActivitiesGroupsServiceFilterReq{})
}

func downloadActivitiesStatusesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewActivitiesStatusesServiceClient(conn).DownloadAsCSV(ctx, &sdk.ActivitiesStatusesServiceFilterReq{})
}

func downloadActivitiesTagsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewActivitiesTagsServiceClient(conn).DownloadAsCSV(ctx, &sdk.ActivitiesTagsServiceFilterReq{})
}

func downloadActivitiesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewActivitiesServiceClient(conn).DownloadAsCSV(ctx, &sdk.ActivitiesServiceFilterReq{})
}

func downloadAnnouncementsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewAnnouncementsServiceClient(conn).DownloadAsCSV(ctx, &sdk.AnnouncementsServiceFilterReq{})
}

func downloadAssociatesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewAssociatesServiceClient(conn).DownloadAsCSV(ctx, &sdk.AssociatesServiceFilterReq{})
}

func downloadBankAccountsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewBankAccountsServiceClient(conn).DownloadAsCSV(ctx, &sdk.BankAccountsServiceFilterReq{})
}

func downloadClientsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewClientsServiceClient(conn).DownloadAsCSV(ctx, &sdk.ClientsServiceFilterReq{})
}

func downloadComponentsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewComponentsServiceClient(conn).DownloadImportTemplate(ctx, &sdk.Empty{})
}

func downloadCurrenciesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewCurrenciesServiceClient(conn).DownloadAsCSV(ctx, &sdk.CurrenciesServiceFilterReq{})
}

func downloadDepartmentsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewDepartmentsServiceClient(conn).DownloadAsCSV(ctx, &sdk.DepartmentsServiceFilterReq{})
}

func downloadEquationsFamiliesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquationsFamiliesServiceClient(conn).DownloadAsCSV(ctx, &sdk.EquationsFamiliesServiceFilterReq{})
}

func downloadEquationsReplaceablesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquationsReplaceablesServiceClient(conn).DownloadAsCSV(ctx, &sdk.EquationsReplaceablesServiceFilterReq{})
}

func downloadEquationsSalesBundlesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquationsSalesBundlesServiceClient(conn).DownloadAsCSV(ctx, &sdk.EquationsSalesBundlesServiceFilterReq{})
}

func downloadEquationsWorkOrdersTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquationsWorkOrdersServiceClient(conn).DownloadAsCSV(ctx, &sdk.EquationsWorkOrdersServiceFilterReq{})
}

func downloadEquipmentsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquipmentsServiceClient(conn).DownloadImportTemplate(ctx, &sdk.Empty{})
}

func downloadFamiliesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewFamiliesServiceClient(conn).DownloadAsCSV(ctx, &sdk.FamiliesServiceFilterReq{})
}

func downloadFeedstocksTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewFeedstocksServiceClient(conn).DownloadImportTemplate(ctx, &sdk.Empty{})
}

func downloadFormsFieldsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewFormsFieldsServiceClient(conn).DownloadAsCSV(ctx, &sdk.FormsFieldsServiceFilterReq{})
}

func downloadFormsSectionsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewFormsSectionsServiceClient(conn).DownloadAsCSV(ctx, &sdk.FormsSectionsServiceFilterReq{})
}

func downloadHolidaysTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewHolidaysServiceClient(conn).DownloadAsCSV(ctx, &sdk.HolidaysServiceFilterReq{})
}

func downloadInfrastructuresTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewInfrastructuresServiceClient(conn).DownloadImportTemplate(ctx, &sdk.Empty{})
}

func downloadLabelsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewLabelsServiceClient(conn).DownloadAsCSV(ctx, &sdk.LabelsServiceFilterReq{})
}

func downloadLeavesTypesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewLeavesTypesServiceClient(conn).DownloadAsCSV(ctx, &sdk.LeavesTypesServiceFilterReq{})
}

func downloadLedgersTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewLedgersServiceClient(conn).DownloadAsCSV(ctx, &sdk.LedgersServiceFilterReq{})
}

func downloadLocationsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewLocationsServiceClient(conn).DownloadAsCSV(ctx, &sdk.LocationsServiceFilterReq{})
}

func downloadMerchandisesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewMerchandisesServiceClient(conn).DownloadImportTemplate(ctx, &sdk.Empty{})
}

func downloadPayrollGroupsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewPayrollGroupsServiceClient(conn).DownloadAsCSV(ctx, &sdk.PayrollGroupsServiceFilterReq{})
}

func downloadPayrollParamsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewPayrollParamsServiceClient(conn).DownloadAsCSV(ctx, &sdk.PayrollParamsServiceFilterReq{})
}

func downloadProductsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewProductsServiceClient(conn).DownloadImportTemplate(ctx, &sdk.Empty{})
}

func downloadQCGroupsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewQCGroupsServiceClient(conn).DownloadAsCSV(ctx, &sdk.QCGroupsServiceFilterReq{})
}

func downloadQCParamsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewQCParamsServiceClient(conn).DownloadAsCSV(ctx, &sdk.QCParamsServiceFilterReq{})
}

func downloadRolesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewRolesServiceClient(conn).DownloadAsCSV(ctx, &sdk.RolesServiceFilterReq{})
}

func downloadShiftsGroupsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewShiftsGroupsServiceClient(conn).DownloadAsCSV(ctx, &sdk.ShiftsGroupsServiceFilterReq{})
}

func downloadShiftsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewShiftsServiceClient(conn).DownloadAsCSV(ctx, &sdk.ShiftsServiceFilterReq{})
}

func downloadSkillsGroupsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewSkillsGroupsServiceClient(conn).DownloadAsCSV(ctx, &sdk.SkillsGroupsServiceFilterReq{})
}

func downloadSkillsParamsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewSkillsParamsServiceClient(conn).DownloadAsCSV(ctx, &sdk.SkillsParamsServiceFilterReq{})
}

func downloadStoragesTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewStoragesServiceClient(conn).DownloadAsCSV(ctx, &sdk.StoragesServiceFilterReq{})
}

func downloadStoresTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewStoresServiceClient(conn).DownloadAsCSV(ctx, &sdk.StoresServiceFilterReq{})
}

func downloadTaxGroupsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewTaxGroupsServiceClient(conn).DownloadAsCSV(ctx, &sdk.TaxGroupsServiceFilterReq{})
}

func downloadTaxParamsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewTaxParamsServiceClient(conn).DownloadAsCSV(ctx, &sdk.TaxParamsServiceFilterReq{})
}

func downloadTeamsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewTeamsServiceClient(conn).DownloadAsCSV(ctx, &sdk.TeamsServiceFilterReq{})
}

func downloadUnitsOfMaterialsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewUnitsOfMaterialsServiceClient(conn).DownloadAsCSV(ctx, &sdk.UnitsOfMaterialsServiceFilterReq{})
}

func downloadUsersTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewUsersServiceClient(conn).DownloadAsCSV(ctx, &sdk.UsersServiceFilterReq{})
}

func downloadVendorsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewVendorsServiceClient(conn).DownloadAsCSV(ctx, &sdk.VendorsServiceFilterReq{})
}

func downloadPrimaryRecords(ctx context.Context, conn *grpc.ClientConn) {
	// Action Codes
	savePrimaryRecordsToFile(downloadActionCodesTemplate(ctx, conn))
	// Activity Groups
	savePrimaryRecordsToFile(downloadActivitiesGroupsTemplate(ctx, conn))
	// Activity Statuses
	savePrimaryRecordsToFile(downloadActivitiesStatusesTemplate(ctx, conn))
	// Activity Tags
	savePrimaryRecordsToFile(downloadActivitiesTagsTemplate(ctx, conn))
	// Activities
	savePrimaryRecordsToFile(downloadActivitiesTemplate(ctx, conn))
	// Announcements
	savePrimaryRecordsToFile(downloadAnnouncementsTemplate(ctx, conn))
	// Associates
	savePrimaryRecordsToFile(downloadAssociatesTemplate(ctx, conn))

	// Bank Accounts
	savePrimaryRecordsToFile(downloadBankAccountsTemplate(ctx, conn))
	// Clients
	savePrimaryRecordsToFile(downloadClientsTemplate(ctx, conn))
	// Components
	savePrimaryRecordsToFile(downloadComponentsTemplate(ctx, conn))
	// Currencies
	savePrimaryRecordsToFile(downloadCurrenciesTemplate(ctx, conn))
	// Departments
	savePrimaryRecordsToFile(downloadDepartmentsTemplate(ctx, conn))

	// Equation Families
	savePrimaryRecordsToFile(downloadEquationsFamiliesTemplate(ctx, conn))
	// Equation Replaceables
	savePrimaryRecordsToFile(downloadEquationsReplaceablesTemplate(ctx, conn))
	// Equation Sales Bundles
	savePrimaryRecordsToFile(downloadEquationsSalesBundlesTemplate(ctx, conn))
	// Equation Work Orders
	savePrimaryRecordsToFile(downloadEquationsWorkOrdersTemplate(ctx, conn))

	// Equipments
	savePrimaryRecordsToFile(downloadEquipmentsTemplate(ctx, conn))
	// Families
	savePrimaryRecordsToFile(downloadFamiliesTemplate(ctx, conn))
	// Feedstocks
	savePrimaryRecordsToFile(downloadFeedstocksTemplate(ctx, conn))
	// Form Fields
	savePrimaryRecordsToFile(downloadFormsFieldsTemplate(ctx, conn))
	// Form Sections
	savePrimaryRecordsToFile(downloadFormsSectionsTemplate(ctx, conn))
	// Holidays
	savePrimaryRecordsToFile(downloadHolidaysTemplate(ctx, conn))
	// Infrastructures
	savePrimaryRecordsToFile(downloadInfrastructuresTemplate(ctx, conn))
	// Labels
	savePrimaryRecordsToFile(downloadLabelsTemplate(ctx, conn))
	// Leave Types
	savePrimaryRecordsToFile(downloadLeavesTypesTemplate(ctx, conn))
	// Ledgers
	savePrimaryRecordsToFile(downloadLedgersTemplate(ctx, conn))
	// Locations
	savePrimaryRecordsToFile(downloadLocationsTemplate(ctx, conn))
	// Merchandises
	savePrimaryRecordsToFile(downloadMerchandisesTemplate(ctx, conn))
	// Payroll Groups
	savePrimaryRecordsToFile(downloadPayrollGroupsTemplate(ctx, conn))
	// Payroll Params
	savePrimaryRecordsToFile(downloadPayrollParamsTemplate(ctx, conn))
	// Products
	savePrimaryRecordsToFile(downloadProductsTemplate(ctx, conn))

	// QC Groups
	savePrimaryRecordsToFile(downloadQCGroupsTemplate(ctx, conn))
	// QC Params
	savePrimaryRecordsToFile(downloadQCParamsTemplate(ctx, conn))

	// Roles
	savePrimaryRecordsToFile(downloadRolesTemplate(ctx, conn))

	// Shift Groups
	savePrimaryRecordsToFile(downloadShiftsGroupsTemplate(ctx, conn))
	// Shifts
	savePrimaryRecordsToFile(downloadShiftsTemplate(ctx, conn))
	// Skill Groups
	savePrimaryRecordsToFile(downloadSkillsGroupsTemplate(ctx, conn))
	// Skill Params
	savePrimaryRecordsToFile(downloadSkillsParamsTemplate(ctx, conn))
	// Storages
	savePrimaryRecordsToFile(downloadStoragesTemplate(ctx, conn))
	// Stores
	savePrimaryRecordsToFile(downloadStoresTemplate(ctx, conn))
	// Tax Groups
	savePrimaryRecordsToFile(downloadTaxGroupsTemplate(ctx, conn))
	// Tax Params
	savePrimaryRecordsToFile(downloadTaxParamsTemplate(ctx, conn))
	// Teams
	savePrimaryRecordsToFile(downloadTeamsTemplate(ctx, conn))
	// Units of Materials
	savePrimaryRecordsToFile(downloadUnitsOfMaterialsTemplate(ctx, conn))
	// Users
	savePrimaryRecordsToFile(downloadUsersTemplate(ctx, conn))
	// Vendors
	savePrimaryRecordsToFile(downloadVendorsTemplate(ctx, conn))

	fmt.Println("Generated all the primary CSV files successfully")
}

func downloadAssetIndentsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewAssetIndentsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadCreditNotesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewCreditNotesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadDebitNotesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewDebitNotesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadEquationsFamiliesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquationsFamiliesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadEquationsReplaceablesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquationsReplaceablesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadEquationsSalesBundlesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquationsSalesBundlesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadEquationsWorkOrdersLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewEquationsWorkOrdersServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadExpensesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewExpensesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadGoalsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewGoalsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadGoodsReceiptsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewGoodsReceiptsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadPayrollGroupsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewPayrollGroupsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadProductionIndentsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewProductionIndentsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadProductionPlansLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewProductionPlansServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadProformaInvoicesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewProformaInvoicesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadPurchaseEnquiriesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewPurchasesEnquiriesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadPurchaseIndentsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewPurchasesIndentsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadPurchaseOrdersLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewPurchasesOrdersServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadQCGroupsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewQCGroupsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadQuotationRequestsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewQuotationsRequestsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadReplaceableIndentsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewReplaceableIndentsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadSalesEnquiriesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewSalesEnquiriesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadSalesInvoicesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewSalesInvoicesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadSalesOrdersLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewSalesOrdersServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadSalesQuotationsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewSalesQuotationsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadSkillGroupsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewSkillsGroupsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadStockAuditsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewStockAuditsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadSupplyOffersLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewSupplyOffersServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadVendorInvoicesLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewVendorInvoicesServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadVendorsLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewVendorsServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadWorkOrdersLineItemsTemplate(ctx context.Context, conn *grpc.ClientConn) (*sdk.StandardFile, error) {
	return sdk.NewWorkOrdersServiceClient(conn).DownloadItemsTemplateAsCSV(ctx, &sdk.Empty{})
}

func downloadLineItems(ctx context.Context, conn *grpc.ClientConn) {
	// Asset Indents
	saveLineItemsToFile(downloadAssetIndentsLineItemsTemplate(ctx, conn))
	// Credit Notes
	saveLineItemsToFile(downloadCreditNotesLineItemsTemplate(ctx, conn))
	// Debit Notes
	saveLineItemsToFile(downloadDebitNotesLineItemsTemplate(ctx, conn))
	// Equations Families
	saveLineItemsToFile(downloadEquationsFamiliesLineItemsTemplate(ctx, conn))
	// Equations Replaceables
	saveLineItemsToFile(downloadEquationsReplaceablesLineItemsTemplate(ctx, conn))
	// Equations Sales Bundles
	saveLineItemsToFile(downloadEquationsSalesBundlesLineItemsTemplate(ctx, conn))
	// Equations Work Orders
	saveLineItemsToFile(downloadEquationsWorkOrdersLineItemsTemplate(ctx, conn))
	// Expenses
	saveLineItemsToFile(downloadExpensesLineItemsTemplate(ctx, conn))
	// Goals
	saveLineItemsToFile(downloadGoalsLineItemsTemplate(ctx, conn))
	// Goods Receipts
	saveLineItemsToFile(downloadGoodsReceiptsLineItemsTemplate(ctx, conn))
	// Payroll Groups
	saveLineItemsToFile(downloadPayrollGroupsLineItemsTemplate(ctx, conn))
	// Production Indents
	saveLineItemsToFile(downloadProductionIndentsLineItemsTemplate(ctx, conn))
	// Production Plans
	saveLineItemsToFile(downloadProductionPlansLineItemsTemplate(ctx, conn))
	// Proforma Invoices
	saveLineItemsToFile(downloadProformaInvoicesLineItemsTemplate(ctx, conn))
	// Purchase Enquiries
	saveLineItemsToFile(downloadPurchaseEnquiriesLineItemsTemplate(ctx, conn))
	// Purchase Indents
	saveLineItemsToFile(downloadPurchaseIndentsLineItemsTemplate(ctx, conn))
	// Purchase Orders
	saveLineItemsToFile(downloadPurchaseOrdersLineItemsTemplate(ctx, conn))
	// QC Groups
	saveLineItemsToFile(downloadQCGroupsLineItemsTemplate(ctx, conn))
	// Quotation Requests
	saveLineItemsToFile(downloadQuotationRequestsLineItemsTemplate(ctx, conn))
	// Replaceable Indents
	saveLineItemsToFile(downloadReplaceableIndentsLineItemsTemplate(ctx, conn))
	// Sales Enquiries
	saveLineItemsToFile(downloadSalesEnquiriesLineItemsTemplate(ctx, conn))
	// Sales Invoices
	saveLineItemsToFile(downloadSalesInvoicesLineItemsTemplate(ctx, conn))
	// Sales Orders
	saveLineItemsToFile(downloadSalesOrdersLineItemsTemplate(ctx, conn))
	// Sales Quotations
	saveLineItemsToFile(downloadSalesQuotationsLineItemsTemplate(ctx, conn))
	// Skill Groups
	saveLineItemsToFile(downloadSkillGroupsLineItemsTemplate(ctx, conn))
	// Stock Audits
	saveLineItemsToFile(downloadStockAuditsLineItemsTemplate(ctx, conn))
	// Supply Offers
	saveLineItemsToFile(downloadSupplyOffersLineItemsTemplate(ctx, conn))
	// Vendor Invoices
	saveLineItemsToFile(downloadVendorInvoicesLineItemsTemplate(ctx, conn))
	// Vendors
	saveLineItemsToFile(downloadVendorsLineItemsTemplate(ctx, conn))
	// Work Orders
	saveLineItemsToFile(downloadWorkOrdersLineItemsTemplate(ctx, conn))

	fmt.Println("Generated all the line items CSV files successfully")
}
