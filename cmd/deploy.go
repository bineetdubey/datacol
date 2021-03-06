package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	pb "github.com/dinesh/datacol/api/models"
	"github.com/dinesh/datacol/cmd/stdcli"
	"gopkg.in/urfave/cli.v2"
)

func init() {
	stdcli.AddCommand(&cli.Command{
		Name:   "deploy",
		Usage:  "deploy an app",
		Action: cmdDeploy,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "image",
				Aliases: []string{"i"},
				Usage:   "docker image to use",
			},
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "service port",
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build id to use",
			},
			&cli.BoolFlag{
				Name:    "wait",
				Aliases: []string{"w"},
				Usage:   "Wait for the app become available",
				Value:   true,
			},
			&cli.StringFlag{
				Name:    "file, f",
				Aliases: []string{"f"},
				Usage:   "path of Dockerfile or app.yaml",
			},
		},
	})
}

func cmdDeploy(c *cli.Context) error {
	dir, name, err := getDirApp(".")
	if err != nil {
		return err
	}

	client, close := getApiClient(c)
	defer close()

	app, err := client.GetApp(name)
	if err != nil {
		log.Warn(err)
		return app404Err(name)
	}

	var build *pb.Build
	buildId := c.String("build")

	if len(buildId) == 0 {
		if build, err = executeBuildDir(client, app, dir); err != nil {
			return err
		}
	} else {
		b, err := client.GetBuild(name, buildId)
		if err != nil {
			return err
		}

		if b == nil {
			return fmt.Errorf("No build found by id: %s.", buildId)
		}

		build = b
	}

	fmt.Printf("Deploying build %s\n", build.Id)

	if _, err = client.ReleaseBuild(build, c.Bool("wait")); err != nil {
		return err
	}

	app, err = client.GetApp(name)
	if err != nil {
		return fmt.Errorf("fetching app %s err: %v", name, err)
	}

	if len(app.Endpoint) > 0 {
		fmt.Printf("\nDeployed at %s\n", app.Endpoint)
	} else {
		fmt.Println("[DONE].")
	}

	return nil
}
