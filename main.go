package main

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
	"github.com/urfave/cli"
)

func init() {
	// Adding this to avoid the error being pushed to stderr for the second time,
	// since the util method already logs the error with rich formatting and fields.
	cli.ErrWriter = ioutil.Discard
}

func main() {

	var subsCred AzureSubscriptionCred
	var storCred AzureStorageCred

	app := cli.NewApp()

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "clientID, cid",
			Destination: &subsCred.ClientID,
			EnvVar:      "AZ_CLIENT_ID",
			Usage: "One of Client IDs in Azure Subscription. " +
				"Refer https://azure.microsoft.com/en-us/documentation/" +
				"articles/resource-group-create-service-principal-portal/",
		},
		cli.StringFlag{
			Name:        "clientSecret, cpwd",
			Destination: &subsCred.ClientSecret,
			EnvVar:      "AZ_CLIENT_SECRET",
			Usage:       "Client Secret corresponding to the Client.",
		},
		cli.StringFlag{
			Name:        "subscriptionID, sid",
			Destination: &subsCred.SubscriptionID,
			EnvVar:      "AZ_SUBSCRIPTION_ID",
			Usage:       "ID of the Azure Subscription, in which the VMs are hosted.",
		},
		cli.StringFlag{
			Name:        "tenantID, tid",
			Destination: &subsCred.TenantID,
			EnvVar:      "AZ_TENANT_ID",
			Usage:       "Tenant ID corresponding to the Client",
		},
		cli.StringFlag{
			Name:        "storAccName, san",
			Destination: &storCred.AccountName,
			EnvVar:      "AZ_STOR_ACC_NAME",
			Usage:       "Azure storage account name.",
		},
		cli.StringFlag{
			Name:        "storAccKey, sak",
			Destination: &storCred.AccountKey,
			EnvVar:      "AZ_STOR_ACC_KEY",
			Usage:       "Azure storage account key.",
		},
		cli.StringFlag{
			Name:        "storContName, cont",
			Destination: &storCred.ContainerName,
			EnvVar:      "AZ_STOR_CONT_NAME",
			Usage:       "Name of Storage Container to store Data disks.",
		},
	}

	app.Before = func(ctxt *cli.Context) error {
		log.Info("azuredatadisk-dockervolumedriver plugin is starting...")
		return nil
	}

	app.Action = func(ctxt *cli.Context) error {

		var driver volume.Driver
		handler := volume.NewHandler(driver)

		log.Infof("Starting to listen to the socket file : %s", UnixSocketFileName)
		if err := handler.ServeUnix("root", UnixSocketFileName); err != nil {
			return makeErrorFromErr(
				ErrorSocketFileFailure,
				err,
				"Failed to listen to the socket file : %s", UnixSocketFileName)
		}

		return nil
	}

	app.After = func(ctxt *cli.Context) error {
		log.Info("azuredatadisk-dockervolumedriver plugin is stopping...")
		return nil
	}

	app.Flags = flags

	app.Run(os.Args)
}
