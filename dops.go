package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/mgutz/ansi"
	"github.com/samalba/dockerclient"
)

func getIPs(es map[string]dockerclient.EndpointSettings) []string {
	s := []string{}
	for _, v := range es {
		s = append(s, v.IPAddress)
	}
	return s
}

func getPorts(ps []dockerclient.Port) []string {
	if len(ps) == 0 {
		s := make([]string, 1)
		s[0] = ""
		return s
	}
	s := make([]string, len(ps))
	for i, p := range ps {
		s[i] = ""
		if p.Type == "udp" {
			s[i] += p.Type + ":"
		}
		if p.IP != "" {
			s[i] += "h/" + strconv.Itoa(p.PublicPort) + "->"
		}
		s[i] += "c/" + strconv.Itoa(p.PrivatePort)
	}
	return s
}

type ColorFunc func(string) string

func first(x int, s string) string {
	if x > len(s) {
		return s[:len(s)]
	} else {
		return s[:x]
	}
}

func getNewColor(m *map[string]struct{}, old string) string {
	// the range for map is random
	var picked string
	for c, _ := range *m {
		picked = c
		break
	}
	delete(*m, picked)
	(*m)[old] = struct{}{}
	return picked
}

func main() {
	sck := os.Getenv("DOCKER_HOST")
	if sck == "" {
		sck = "unix:///var/run/docker.sock"
	}
	client, err := dockerclient.NewDockerClient(sck, nil)
	if err != nil {
		log.Fatal(err)
	}

	containers, err := client.ListContainers(false, false, "")
	if err != nil {
		log.Fatal(err)
	}

	w := new(tabwriter.Writer)
	usedColor := "red"
	w.Init(os.Stdout, 0, 2, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tImage\tIP\tPorts\tName")
	other_colors := map[string]struct{}{
		"blue":  struct{}{},
		"green": struct{}{},
		"cyan":  struct{}{}}

	for _, c := range containers {
		newColor := getNewColor(&other_colors, usedColor)
		colorize := ansi.ColorFunc(newColor)
		cc, err := client.InspectContainer(c.Id)
		if err != nil {
			log.Fatal(err)
		}
		portlines := getPorts(c.Ports)
		ips := getIPs(c.NetworkSettings.Networks)
		line := c.Id[:4] + "\t" + first(25, c.Image) + "\t" + ips[0] + "\t" + portlines[0] + "\t" + first(20, cc.Name)

		fmt.Fprintln(w, colorize(line))
		if len(portlines) > 1 {
			for _, l := range portlines[1:] {
				fmt.Fprintln(w, colorize(" \t \t \t"+l+"\t "))
			}
		}
		usedColor = newColor

	}
	w.Flush()
}
