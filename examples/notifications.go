package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	 "github.com/HewlettPackard/oneview-golang/utils"
//	"os"
//	"strconv"
)

func main() {
	var (
		ClientOV    *ov.OVClient
	)
	//apiversion, _ := strconv.Atoi(os.Getenv("1600"))
	ovc := ClientOV.NewOVClient(
		"Administrator",
	        "admin123",
		"LOCAL",
		"https://10.50.9.90/",
		false,
		1600,
		"*")

	// Retrieve email notification details with configured filters
	emailNotifications, err := ovc.GetEmailNotifications("", "", "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get Email Notifications----#")
	 	fmt.Println(emailNotifications)
		
	}

	// Retrieve email notification filter names for a given scope
	filter := "scope:scopename1"
	emailNotificationsByFilter, err := ovc.GetEmailNotificationsByFilter(filter, "", "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get Email Notifications by Filter----------------#")
		fmt.Println(emailNotificationsByFilter)
	}

	// Retrieve test email notification configuration details
	emailNotificationConfiguration, err := ovc.GetEmailNotificationsConfiguration("", "", "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get Email Notifications by Configuration----------------#")
		fmt.Println(emailNotificationConfiguration)
	}

	// Sends test email from appliance to specified user(s).
	toAddress := &[]utils.Nstring{utils.NewNstring("email1@example.com")}
	testEmail := ov.TestEmailRequest{
		HtmlMessageBody:         "Html alert message with html and css content",
		Subject:                 "Critical alert generated",
		TextMessageBody:         "Plain text mail content",
		ToAddress:               *toAddress,
	  },

	  err := ovc.SendTestEmail(testEmail)
	  if err != nil {
		fmt.Println("Sending TestEmail Failed: ", err)
	  } else {
		fmt.Println("TestEmail sent successfully...")
	  }

	// Sends email from appliance to specified user
//	email = ov.TestEmailRequest{
//		htmlMessageBody:         "Html alert message with html and css content",
//		subject:                 "Critical alert generated",
//		textMessageBody:         "Plain text mail content",
//		toAddress:               "email1@example.com",
//	  }
//	err = ovc.SendEmail(email)
//	if err != nil {
//		fmt.Println("Sending Email Failed: ", err)
//	} else {
//		fmt.Println("Email sent successfully...")
//	}

	// Configure the appliance to send an email notification, generated by specified alert filter queries.
//	alertEmailFilter1 = ov.AlertEmailFilters{disabled: false,
//		displayFilter:            "status:warning status:critical",
//		emails:                   "email1@example.com",
//		filter:                   "status:'warning' or  status:'critical'",
//		userQueryFilter:          "CPU",
//		limit:                     3,
//		limitDuration:             "minute",
///		scopeQuery:                "scope:'windows'",
//		filterName:                "Critical Alerts",
//	},
//	alertEmailFilters = new([]ov.AlertEmailFilters)
//	*alertEmailFilters = append(*alertEmailFilters, alertEmailFilter1)
//	email =  ov.EmailNotification{
//		alertEmailDisabled:       false,
//		alertEmailFilters:        *alertEmailFilters,
//		password:                 "some password",
//		senderEmailAddress:       "sender@example.com",
///		smtpPort:                 25,
//		smtpServer:               "smtp.example.com",
//		smtpProtocol:             "TLS",
//	  },
//	err2 = ovc.ConfigureAppliance(email)
//	if err2 != nil {
///		fmt.Println("Sending Email Failed: ", err2)
//	} else {
//		fmt.Println("Email sent successfully...")
//	}
}
